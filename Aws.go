package faws

import (
	"faws/api/aws"
	"faws/api/ec2"
)

type Amazon struct {
	Auth aws.Auth
	EC2  *ec2.EC2
}

func CreateInstancesOptions() ec2.RunInstances {
	insts := ec2.RunInstances{}
	return insts
}

func New(access_key, access_secret string, region aws.Region) Amazon {
	var auth aws.Auth
	auth.AccessKey = access_key
	auth.SecretKey = access_secret

	//Create EC2 object, for specific region
	e := ec2.New(auth, region)
	am := Amazon{
		Auth: auth,
		EC2:  e,
	}
	return am
}

func GetRegions() map[string]aws.Region {
	return aws.Regions
}

func (amz *Amazon) GetInstances(ids []string, filter string) (instances []ec2.Instance, ret_error error) {
	ret_error = nil
	//TODO: Need to implement filter String
	//Something like "arcitecture: i386|name: ddd"

	fil := ec2.NewFilter()
	resp, err := amz.EC2.Instances(ids, fil)
	if err != nil {
		ret_error = err
		return
	}

	for _, reservations := range resp.Reservations {
		for _, inst := range reservations.Instances {
			instances = append(instances, inst)
		}
	}
	return
}

func (amz *Amazon) StopInstance(id string) (ec2.InstanceState, error) {
	resp, err := amz.EC2.StopInstances(id)
	if err != nil {
		return ec2.InstanceState{}, err
	}

	/*
		|	0    |    pending    |
		|  16    |    running    |
		|  32    | shutting-down |
		|  48    |  terminated   |
		|  64    |   stopping    |
		|  80    |   stopped
	*/
	return resp.StateChanges[0].CurrentState, nil
}

func (amz *Amazon) StartInstance(id string) (ec2.InstanceState, error) {
	resp, err := amz.EC2.StartInstances(id)
	if err != nil {
		return ec2.InstanceState{}, err
	}
	/*
		|	0    |    pending    |
		|  16    |    running    |
		|  32    | shutting-down |
		|  48    |  terminated   |
		|  64    |   stopping    |
		|  80    |   stopped
	*/

	return resp.StateChanges[0].CurrentState, nil
}

func (amz *Amazon) GetImages(ids []string, filter string) (images []ec2.Image, ret_error error) {
	ret_error = nil
	//TODO: Need to implement filter String
	//Something like "arcitecture: i386|name: ddd"

	fil := ec2.NewFilter()
	fil.Add("is-public", "false")
	resp, err := amz.EC2.Images(ids, fil)
	if err != nil {
		ret_error = err
		return
	}

	for _, img := range resp.Images {
		images = append(images, img)
	}
	return
}

func (amz *Amazon) CreateInstance(options ec2.RunInstances) (instances []ec2.Instance, ret_error error) {
	ret_error = nil
	resp, err := amz.EC2.RunInstances(&options)
	if err != nil {
		ret_error = err
		return
	}
	for _, inst := range resp.Instances {
		instances = append(instances, inst)
	}
	return
}

func (amz *Amazon) DeleteInstance(id string) (status ec2.InstanceState, ret_err error) {
	ids := make([]string, 1)
	ids[0] = id
	resp, err := amz.EC2.TerminateInstances(ids)
	if err != nil {
		ret_err = err
		return
	}
	status = resp.StateChanges[0].CurrentState
	return
}
