package sign

import (
	"fmt"
	"testing"
)

func TestVerify_ParseQuery(t *testing.T) {
	requestURI := "/restful/api/numbers?app_id=9d8a121ce581499d&nonce_str=tempstring&city=beijing" +
		"&timestamp=1532585241&sign=0f5b8c97920bc95f1a8b893f41b42d9e"

	// 第一步：创建Verify校验类
	verifier := NewVerifier()

	// 假定从RequestUri中读取校验参数
	if err := verifier.ParseQuery(requestURI); nil != err {
		t.Fatal(err)
	}

	// 第二步：（可选）校验是否包含签名校验必要的参数
	if err := verifier.MustHasOtherKeys("city"); nil != err {
		t.Fatal(err)
	}

	// 第三步：检查时间戳是否超时。
	//if err := verifier.CheckTimeStamp(); nil != err {
	//	t.Fatal(err)
	//}

	// 第四步，创建Sign来重现客户端的签名信息：
	signer := NewSignerMd5()

	// 第五步：从Verify中读取所有请求参数
	signer.SetBody(verifier.GetBodyWithoutSign())

	// 第六步：从数据库读取AppID对应的SecretKey
	// appId := verifier.MustString("app_id")
	secretKey := "d93047a4d6fe6111"

	// 使用同样的WrapBody方式
	signer.SetAppSecretWrapBody(secretKey)

	// 生成
	sign := signer.GetSignature()
	t.Log("sign", sign)

	// 校验自己生成的和传递过来的是否一致
	if verifier.MustString("sign") != sign {
		t.Fatal("校验失败")
	}

	fmt.Println(sign)
}

func TestVerify_Rsa(t *testing.T){
	requestURI := "http://34.56.56.55?app_id=112233&nonce_str=supertempstr&tags=github,gopher&timestamp=1594458195&username=1024casts&sign=6a44625938387374337045505f6a693537466177706c4b63634a6379734b5636535f444b45386e4b4d7764514b305471427a4c7454374c575a724d454f65624d6d5258565141326e6a786b70677a36776e457334315376724179624e6c7658787056644f59706b6d766c66464f44505033324e435f492d4b672d776530735764486933737336773559766373675a4c5a7054686735347a707236455570375042533655335067586f525a61467a73372d555f65596c38745542574e6d6750786d6b34394d4f6b544d704851486b414f37676663453279666b2d386c52366a747a6d6d35635346684d62696364584b644c37354c716258594e6539593631736567564d73634d615145437a7766514765334c3551776a5f5958306157637a6f3144615048636141534d75624441776159315f365977676e5a38734d5a554764666d4b4a696137786f484739506939573750433359366276465557695053546e394c6f63355f765a5561726a5f47784c634339474366774a2d7737554a5976355365485339387a377832454c7a51486c4d706a326932305a74384949474d68557279525f77366c425873457343627749487a7a544e347846585a467a556d787933583437416b36505941754967417379494e3654584b3656477265786f6f765f38693045724d6267766855484359454a7152643575657a686e41586759667048325f6d69675f505a592d6471527646684879464846465f71554d6e35457a634d3676786b696a387573636f445f6b65496f59695138585155794c6b78374959494d7765587637747661356b68776365672d6f423933487230536b6e754d32326868382d4c506c7a3047304b4a4c456f3757646c73367a364e6f7567486f42426e694f66396f616271694e2d4b57555249304b426d5535537237494662527934665a326b5a57565150366a6d4d6b3d"

	// 第一步：创建Verify校验类
	verifier := NewVerifier()

	if err := verifier.ParseQuery(requestURI); nil != err {
		t.Fatal(err)
	}
	//if err := verifier.CheckTimeStamp(); nil != err {
	//	t.Fatal(err)
	//}
	// 第四步，创建Sign来重现客户端的签名信息：
	sign := verifier.GetSign()

	println(sign)
	decryptResult := RsaDecryptSign(sign,"/Users/临时/rsa_pri_key")
	fmt.Println("输出签名：" + string(decryptResult))

	//sign := signer.GetSignature()
	//t.Log("sign", sign)
	//
	//// 校验自己生成的和传递过来的是否一致
	//if verifier.MustString("sign") != sign {
	//	t.Fatal("校验失败")
	//}
	//
	//fmt.Println(sign)

}