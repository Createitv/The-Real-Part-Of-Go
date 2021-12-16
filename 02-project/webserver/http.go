package main

/*
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)  处理路由返回函数的接口
}
*/

/*
The HandlerFunc type is an adapter to allow the use of
ordinary functions as HTTP handlers.
type HandlerFunc func(ResponseWriter, *Request)
实现ServeHTTP方法
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
*/
/*
http.HandlerFunc是一个路由接口pattern传入一个string， 第二个参数是处理路由的返回函数
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
*/

//func main() {
//	http.HandleFunc("/", indexHandler)
//	http.HandleFunc("/hello", helloHandler)
//	log.Fatal(http.ListenAndServe(":9999", nil))
//}
//
//func helloHandler(w http.ResponseWriter, req *http.Request) {
//	for k, v := range req.Header {
//		fmt.Printf("Header[%q] = %q\n", k, v)
//	}
//}
//
//func indexHandler(w http.ResponseWriter, request *http.Request) {
//	fmt.Fprintf(w, "URL.Path: %q\n", request.URL.Path)
//}
