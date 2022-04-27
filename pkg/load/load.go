/**
 * @author jiangshangfang
 * @date 2022/3/4 11:22 PM
 **/
package load

type Load interface {
	//加载配置
	LoadConfiguration(name string) (string, error)

	//监听配置变化
	ListenConfiguration(name string) (string, error)
}
