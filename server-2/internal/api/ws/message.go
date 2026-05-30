package ws

type Message struct {
	Type       string `json:"type"`
	HospitalID string `json:"hospitalId,omitempty"`

	RobotID string `json:"robotId,omitempty"`

	// Telepresence (WebRTC)
	SDP       string `json:"sdp,omitempty"`
	Candidate string `json:"candidate,omitempty"`

	// Optional payload (vitals or others)
	Data string `json:"data,omitempty"`
}