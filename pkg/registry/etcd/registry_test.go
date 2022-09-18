/**
 * @author jiangshangfang
 * @date 2022/1/13 7:46 PM
 **/
package etcd

import (
	"context"
	"fmt"
	"gin-micro-mine/pkg/registry"
	"go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestRegistry(t *testing.T) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:12379", "127.0.0.1:22379", "127.0.0.1:32379"},
		DialTimeout: time.Second,
		//DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()
	s := &registry.ServiceInstance{
		ID:      "0",
		Name:    "jonnyClient1",
		Version: "1.1",
	}
	r := New(client)

	if err1 := r.Register(ctx, s); err1 != nil {
		t.Fatal(err1)
	}
	time.Sleep(time.Second)

	res, err := r.GetService(ctx, s.Name)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 1 && res[0].Name != s.Name {
		t.Errorf("not expected: %+v", res)
	}

	w, err := r.Watch(ctx, s.Name)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = w.Stop()
	}()

	go func() {
		for {
			res, err1 := w.Next()
			if err1 != nil {
				return
			}
			t.Logf("watch: %d", len(res))
			for _, r := range res {
				t.Logf("next: %+v", r)
			}
		}
	}()
	//time.Sleep(time.Second)

	//if err1 := r.DeRegister(ctx, s); err1 != nil {
	//	t.Fatal(err1)
	//}

	//res, err = r.GetService(ctx, s.Name)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//if len(res) != 0 {
	//	t.Errorf("not expected empty")
	//}
}

func TestHeartBeat(t *testing.T) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:12379"},
		DialTimeout: time.Second, DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()
	s := &registry.ServiceInstance{
		ID:   "0",
		Name: "jonnyClient2",
	}

	go func() {
		r := New(client)
		w, err1 := r.Watch(ctx, s.Name)
		if err1 != nil {
			return
		}
		defer func() {
			_ = w.Stop()
		}()
		for {
			res, err2 := w.Next()
			if err2 != nil {
				return
			}
			t.Logf("watch: %d", len(res))
			for _, r := range res {
				t.Logf("next: %+v", r)
			}
		}
	}()
	time.Sleep(time.Second)

	// new a server
	r := New(client,
		WithRegisterTTL(2*time.Second),
		WithMaxRetry(5),
	)

	key := fmt.Sprintf("%s/%s/%s", r.opts.namespace, s.Name, s.ID)
	t.Logf(key)
	value, _ := marshal(s)
	leaseID, err := r.registerWithKV(ctx, key, value)
	if err != nil {
		t.Fatal(err)
	}

	// wait for lease expired
	time.Sleep(3 * time.Second)

	res, err := r.GetService(ctx, s.Name)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 0 {
		t.Errorf("not expected empty")
	}

	go r.heartBeat(ctx, leaseID, key, value)

	time.Sleep(time.Second)
	res, err = r.GetService(ctx, s.Name)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) == 0 {
		t.Errorf("reconnect failed")
	}
}

func TestPutAndGet(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:12379"},
		DialTimeout: time.Second, DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "q1mi", "dsb")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "q1mi")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}

func TestWatch(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:12379"},
		DialTimeout: time.Second, DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// watch key:q1mi change
	rch := cli.Watch(context.Background(), "q1mi") // <-chan WatchResponse
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

func TestLease(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:12379"},
		DialTimeout: time.Second, DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	// 创建一个5秒的租约
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		t.Fatal(err)
	}
	//time.Sleep(6*time.Second)
	// 5秒钟之后, /nazha/ 这个key就会被移除
	_, err = cli.Put(context.TODO(), "/nazha/", "dsb", clientv3.WithLease(resp.ID))
	if err != nil {
		t.Fatal(err)
	}
}

func TestKeepAlive(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:12379"},
		DialTimeout: time.Second, DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	// 创建一个5秒的租约
	resp, err := cli.Grant(context.TODO(), 4)
	if err != nil {
		t.Fatal(err)
	}
	//time.Sleep(6*time.Second)
	// 5秒钟之后, /nazha/ 这个key就会被移除
	_, err = cli.Put(context.TODO(), "/nazha/", "dsb", clientv3.WithLease(resp.ID))
	if err != nil {
		t.Fatal(err)
	}

	ch, kaerr := cli.KeepAlive(context.TODO(), resp.ID)
	if kaerr != nil {
		t.Fatal(kaerr)
	}
	for {
		ka := <-ch
		fmt.Println("ttl:", ka.TTL)
	}
}
