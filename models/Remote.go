package models

import (
	"golang.org/x/crypto/ssh"
	"bytes"
	"net"
	"github.com/astaxie/beego/logs"
	"strconv"
)

type Remote struct {
	User string
	Pwd  string
	Add  string
	Port string
}

func (c *Remote)GetSshClientByIp() (*ssh.Client) {
	config := &ssh.ClientConfig{
		User: c.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(c.Pwd),
		},
		//需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	client, err := ssh.Dial("tcp", c.Add + ":" + c.Port, config)
	if err != nil {
		logs.Error("Failed to dial: ", err)
		panic(err)
	}
	return client
}

func (c * Remote)GetSshClientByExec(user string,pwd string,addr string,port int,cmd string) (string, error) {
	c.User = user
	c.Pwd = pwd
	c.Add = addr
	c.Port = strconv.Itoa(port)
	client := c.GetSshClientByIp()
	var buf bytes.Buffer
	session, err := client.NewSession()
	if nil != err {
		return "Failed to NewSession", err
	}
	session.Stdout = &buf
	session.Stderr = &buf
	err = session.Run(cmd)
	if err != nil {
		return "命令执行错误", err
	}
	defer session.Close()
	stdout := buf.String()
	return stdout, nil
}

func (c * Remote)GetSshClientByCMD(user string,pwd string,addr string,port int,cmd string) (string, error) {
	c.User = user
	c.Pwd = pwd
	c.Add = addr
	c.Port = strconv.Itoa(port)
	client := c.GetSshClientByIp()
	session, err := client.NewSession()
	if nil != err {
		return "Failed to NewSession", err
	}
	Stdout,err := session.Output(cmd)
	if err != nil {
		return "命令执行错误", err
	}
	defer session.Close()
    stdout :=string(Stdout)
	return stdout, nil
}
