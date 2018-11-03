package models

type Fire struct {
	Size string
	Temp string
}

type Dog struct {}

func (d *Dog) Cry() string{
	return "わんわん"
}

func (d *Dog) GetFire() Fire{
	return Fire{
		Size: "5",
		Temp: "23度",
	}
}

type Cat struct {}


func (d *Cat) Cry() string{
	return "にゃんにゃん"
}

func (d *Cat) GetFire() Fire{
	return Fire{
		Size: "30",
		Temp: "19度",
	}
}