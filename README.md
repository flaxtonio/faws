# Flaxton Amazon Web Services(AWS) API
<p>
This repository is an implementation of AWS API which is used for <a href="http://flaxton.io">flaxton.io</a> Load Balancer writtent in Go programming language <a href="http://golang.org">golang.org</a>. <i>Some part of code implemented using <a href="https://gopkg.in/amz.v1">gopkg.in/amz.v1</a> libraray</i>
</p>

# Get Started
<p>
<i>For this library you will need to install and learn basics of Go programming language <a href="http://golang.org">golang.org</a></i>.
</p>
<b>
To Get Started you need to follow this few steps.
</b>
<ol>
<li>
Create AWS API key and secret by following Amazon documentation: <a href="http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSGettingStartedGuide/AWSCredentials.html">AWS Credentials</a>
</li>
<li>
Put this library in your project path (GOPATH) using
<code>git clone https://github.com/flaxtonio/faws.git</code>
</li>
<li>
Import this library to you Go language project file 
<pre>
import (
    .
    .
    "faws"
    .
    .
)
</pre>
</li>
</ol>

# "Hello World"
```go
package main

import (
	"faws"
	"faws/api/aws"
	"fmt"
)

func main() {
	amz := faws.New("Your Key", "Your Secret", aws.USWest2)
	instances, err := amz.GetInstances(nil, "")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Your Instances")
	for _, inst := range instances {
		fmt.Println(inst.InstanceId, inst.IPAddress)
	}
}
```
<b>Read <a href="https://github.com/flaxtonio/faws/blob/master/tests/awsTest.go" target="_blank"><code>tests/awsTest.go</code></a> file for more detailed example</b>

# API Functions
<ul>
<li><code>faws.New("Your Key", "Your Secret", AWS_Region)</code> returns object for faws structure</li>
<li><code>faws.GetRegions()</code> returns AWS available regions as an object</li>
<li><code>(amz *Amazon) GetInstances(ids []string, filter string)</code> returns all information about specified instances. If <code>ids</code> is <code>nil</code> function returns all available instances</li>
<li><code>(amz *Amazon) GetImages(ids []string, filter string)</code> returns all information about specified AMI Images. If <code>ids</code> is <code>nil</code> function returns all available AMI Images</li>
<li><code>(amz *Amazon) CreateInstance(options ec2.RunInstances)</code> Cerates an AWS EC2 instance using specified options from ec2.RunInstances structure</li>
<li><code>(amz *Amazon) DeleteInstance(id string)</code> Deletes Instance from AWS account by Termination EC2 instance</li>
<li><code>(amz *Amazon) StartInstance(id string)</code> Starts EC2 instance with specified ID</li>
<li><code>(amz *Amazon) StopInstance(id string)</code> Sending Stop command for EC2 instance</li>
</ul>
