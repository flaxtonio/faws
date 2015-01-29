package main

import (
	"bufio"
	"faws"
	"faws/api/aws"
	"fmt"
	"os"
	"strings"
)

func main() {
	amz := faws.New("Your Key", "Your Secret", aws.USWest2)
	instances, err1 := amz.GetInstances(nil, "")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println("Instances")
	for _, inst := range instances {
		fmt.Println(inst.InstanceId, inst.IPAddress)
	}
	images, err2 := amz.GetImages(nil, "")
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}
	fmt.Println("Images")
	for _, img := range images {
		fmt.Println(img.Id, img.Name)
	}

	for {

		fmt.Println("\n", "Command[start, stop, state, delete, create] InstanceID")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		cmd := strings.Split(text, " ")
		switch cmd[0] {
		case "start":
			{
				status, err := amz.StartInstance(cmd[1])
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(status.Name)
			}
		case "stop":
			{
				status, err := amz.StopInstance(cmd[1])
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(status.Name)
			}
		case "state":
			{
				img, err := amz.GetInstances([]string{cmd[1]}, "")
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(img[0].State.Name)
			}
		case "delete":
			{
				st, err := amz.DeleteInstance(cmd[1])
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(st.Name)
			}
		case "create":
			{
				options := faws.CreateInstancesOptions()
				options.ImageId = cmd[1]
				options.InstanceType = "t2.micro"
				insts, e := amz.CreateInstance(options)
				if e != nil {
					fmt.Println(e.Error())
					return
				}
				for _, i := range insts {
					fmt.Println(i.InstanceId)
				}
			}
		}
	}
}
