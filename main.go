package main

/*
func main() {
	//test()
	Listener, err := net.InitKCP("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	for true {
		Conn, err := Listener.AcceptKCP()
		if err != nil {
			continue
		}
		go func(Conn *kcp.UDPSession) {
			str, err := net.GetData(Conn)
			if err != nil {
				panic(err)
			}
			RotaicXBasicTool.Rlog.Println(str)
			Conn.Write(net.MakeData("OK"))
		}(Conn)
	}
}

func test() {
	fmt.Println(net.MakeData("ABC"))
}
*/
/*
func main() {
	privateKey, _ := crypto.NewKey4096()
	bprk := crypto.PrivateKeyToBase64PrivateKey(privateKey)
	fmt.Println(bprk)
	bpuk := crypto.PublicKeyToBase64PublicKey(&privateKey.PublicKey)
	fmt.Println(bpuk)
	prk, _ := crypto.Base64PrivateKeyToPrivateKey(bprk)
	puk, _ := crypto.Base64PublicKeyToPublicKey(bpuk)
	fmt.Println(prk)
	fmt.Println(puk)
}
*/
