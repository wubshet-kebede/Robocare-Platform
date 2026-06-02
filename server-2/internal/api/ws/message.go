package ws

type Message struct {
	Type       string        `json:"type"`
	RobotID    string        `json:"robotId"`
	HospitalID string        `json:"hospitalId,omitempty"`
	SDP        string        `json:"sdp,omitempty"`
	Candidate  *IceCandidate `json:"candidate,omitempty"` 
}
type IceCandidate struct {
	Candidate     string `json:"candidate"`
	SdpMid        string `json:"sdpMid"`
	SdpMLineIndex int    `json:"sdpMLineIndex"`
}