package email

import (
	"ToDoList/internal/utils"
	"crypto/tls"
	"fmt"
	"net/smtp"
)

// EmailManager 邮件系统
type EmailManager struct {
}

// EmailConfig 邮件配置结构体
type EmailConfig struct {
	SMTPHost    string
	SMTPPort    string
	SenderEmail string
	SenderPass  string
}

// Email 邮件结构体
type Email struct {
	To          string
	Subject     string
	Body        string
	Attachments map[string][]byte
}

// newEmail 创建一个新邮件
func newEmail(to string) (Email, error) {
	var coh utils.VerificationCodeManager
	coh = utils.NewVerificationCodeHandler()
	code, err := coh.NewVerificationCode(to)
	if err != nil {
		return Email{}, err
	}
	return Email{
		To:      to,
		Subject: "验证码",
		Body:    "验证码为：\n" + code.Code,
	}, nil
}

// NewEmailManager 创建一个新的邮件管理系统
func NewEmailManager() EmailManager {
	return EmailManager{}
}

// SendSimpleEmail 发送简单文本邮件
func SendSimpleEmail(config EmailConfig, email Email) error {
	// 设置邮件内容
	header := make(map[string]string)
	header["From"] = config.SenderEmail
	header["To"] = email.To
	header["Subject"] = email.Subject
	header["Content-Type"] = "text/plain; charset=UTF-8"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + email.Body

	// 配置 TLS
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         config.SMTPHost,
	}

	// 建立连接
	addr := fmt.Sprintf("%s:%s", config.SMTPHost, config.SMTPPort)
	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return fmt.Errorf("连接到邮件服务器失败: %v", err)
	}
	defer conn.Close()

	// 创建新的 SMTP 客户端
	c, err := smtp.NewClient(conn, config.SMTPHost)
	if err != nil {
		return fmt.Errorf("创建SMTP客户端失败: %v", err)
	}
	defer c.Close()

	// 认证
	auth := smtp.PlainAuth("", config.SenderEmail, config.SenderPass, config.SMTPHost)
	if err = c.Auth(auth); err != nil {
		return fmt.Errorf("SMTP认证失败: %v", err)
	}

	// 设置发件人
	if err = c.Mail(config.SenderEmail); err != nil {
		return fmt.Errorf("设置发件人失败: %v", err)
	}

	// 设置收件人
	if err = c.Rcpt(email.To); err != nil {
		return fmt.Errorf("设置收件人失败: %v", err)
	}

	// 发送邮件内容
	w, err := c.Data()
	if err != nil {
		return fmt.Errorf("创建数据写入器失败: %v", err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("写入邮件内容失败: %v", err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("关闭写入器失败: %v", err)
	}

	return c.Quit()
}

// ConfigureEmail 配置并发送邮件
func (em EmailManager) ConfigureEmail(to string) error {
	config := EmailConfig{
		SMTPHost:    "smtp.163.com",
		SMTPPort:    "465",
		SenderEmail: "m97688596@163.com",
		SenderPass:  "GZMbM37v8M7EaC3P", // 使用授权码
	}

	email, err := newEmail(to)
	if err != nil {
		return err
	}
	// 发送邮件
	err = SendSimpleEmail(config, email)
	if err != nil {
		return err
	}
	return nil
}
