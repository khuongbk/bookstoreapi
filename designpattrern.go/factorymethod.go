package main

import "fmt"

// this is my concrete girl friend
type GirlFriend struct {
	nationality string
	eyesColor   string
	language    string
}

// abstract factory for creating girlfriend
type AbstractFactory interface {
	CreateMyLove() GirlFriend
}

// my indian girlfriend
type IndianGirlFriendFactory struct {
}

// my korean girlfirend
type KoreanGirlFriendFactory struct {
}

// concrete implementation
func (a IndianGirlFriendFactory) CreateMyLove() GirlFriend {
	return GirlFriend{"Indian", "Black", "Hindi"}
}

// concrete implementation
func (a KoreanGirlFriendFactory) CreateMyLove() GirlFriend {
	return GirlFriend{"Korean", "Brown", "Korean"}
}

// main factory method
func getGirlFriend(typeGf string) GirlFriend {

	var gffact AbstractFactory
	switch typeGf {
	case "Indian":
		gffact = IndianGirlFriendFactory{}
		return gffact.CreateMyLove()
	case "Korean":
		gffact = KoreanGirlFriendFactory{}
		return gffact.CreateMyLove()
	}
	return GirlFriend{}
}

func main() {

	var typeGf string

	fmt.Scanf("%s ", &typeGf)
	a := getGirlFriend(typeGf)

	fmt.Println(a)

}

/* pattern factory: tuong tuong la mot nha may, mo hinh co rat nhieu nhung day chuyen(struct).
Nha may nay thi co co dinh cac phuong phap san xuat ma day chuyen nao cung phai lam va duoc declare trong phan  interface
phai specify khi ap dung cac method do len moi day chuyen (struct)thi tao ra duoc cai gi.
Tren day la model cua factory
co main factory method : voi mot dau vao bat ki (bien , doi tuong) thi ko phai dinh nghia trk cai gi ma cu ap dung len main method thi se
cho ra duoc san pham khi ap dung cac phuong phap cua nha may (duoc dinh nghia trong interface)
*/
