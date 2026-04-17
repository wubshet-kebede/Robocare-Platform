package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

type EmailRequest struct {
	To      string
	Subject string
	Body    string
}

func SendInvitationEmail(toEmail, token, role, hospitalName string) error {
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	frontendURL := os.Getenv("FRONTEND_URL")

	if user == "" || pass == "" || host == "" || port == "" {
		return fmt.Errorf("mailer: SMTP credentials are not configured properly")
	}

	inviteLink := fmt.Sprintf("%s/accept-invite?token=%s", frontendURL, token)


	data := struct {
		Role         string
		HospitalName string
		InviteLink   string
	}{
		Role:         role,
		HospitalName: hospitalName,
		InviteLink:   inviteLink,
	}

	
	const emailTemplate = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
</head>
<body style="margin:0;padding:0;background:#f4f6f8;font-family:Arial,sans-serif;">
	<div style="max-width:600px;margin:40px auto;background:#ffffff;border-radius:12px;overflow:hidden;border:1px solid #e5e7eb;">

		<div style="background:#2563eb;padding:20px;text-align:center;color:#fff;">
			<h2 style="margin:0;">{{.HospitalName}}</h2>
			<p style="margin:5px 0 0;font-size:14px;">Staff Invitation</p>
		</div>

		<div style="padding:30px;">
			<h3 style="margin-bottom:10px;color:#111827;">You're invited 🎉</h3>

			<p style="color:#374151;font-size:14px;line-height:1.6;">
				You have been invited to join <b>{{.HospitalName}}</b> as a
				<b style="color:#2563eb;">{{.Role}}</b>.
			</p>

			<div style="text-align:center;margin:30px 0;">
				<a href="{{.InviteLink}}"
				   style="background:#2563eb;color:#fff;text-decoration:none;
				   padding:12px 24px;border-radius:8px;font-weight:600;display:inline-block;">
					Accept Invitation
				</a>
			</div>

			<p style="font-size:12px;color:#6b7280;line-height:1.5;">
				If you were not expecting this invitation, you can safely ignore this email.
				This link will expire automatically for security reasons.
			</p>
		</div>

		<div style="background:#f9fafb;padding:15px;text-align:center;font-size:12px;color:#9ca3af;">
			© 2026 {{.HospitalName}}. All rights reserved.
		</div>

	</div>
</body>
</html>`

	tmpl, err := template.New("invite").Parse(emailTemplate)
	if err != nil {
		return fmt.Errorf("mailer: failed to parse template: %w", err)
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("mailer: failed to execute template: %w", err)
	}

	
	fromHeader := fmt.Sprintf("From: %s <%s>\r\n", hospitalName, user)
	toHeader := fmt.Sprintf("To: %s\r\n", toEmail)
	subjectHeader := "Subject: Hospital Staff Invitation\r\n"
	mime := "MIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n"

	message := []byte(fromHeader + toHeader + subjectHeader + mime + body.String())

	
	auth := smtp.PlainAuth("", user, pass, host)

	addr := fmt.Sprintf("%s:%s", host, port)

	
	err = smtp.SendMail(
		addr,
		auth,
		user,
		[]string{toEmail},
		message,
	)

	if err != nil {
		return fmt.Errorf("mailer: failed to send email via SMTP: %w", err)
	}

	return nil
}