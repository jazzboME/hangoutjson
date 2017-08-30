package hangoutjson

type ID struct {
	ID			string		`json:"ID"`
}

type ResponseHeader struct {
	Status			string		`json:"status"`
	DebugURL		string		`json:"debug_url"`
	RequestTraceID		string		`json:"request_trace_id"`
	CurrentServerTime	string		`json:"current_server_time"`
	BuildLabel		string		`json:"build_label"`
	ChangelistNumber	int		`json:"changelist_number"`
}

type ParticipantID struct {
	GaiaID			string		`json:"gaia_id"`
	ChatID			string		`json:"chat_id"`
}

type SelfReadState struct {
	ParticipantID		ID		`json:"participant_id"`
	LatestReadTimestamp	string		`json:"latest_read_timestamp"`
}

type DeliveryMedium struct {
	MediumType		string		`json:"medium_type"`
}

type DeliveryMediumOption []struct {
	DeliveryMedium		DeliveryMedium	`json:"delivery_medium"`
	CurrentDefault		bool		`json:"current_default"`
}
type SelfConversationState struct {
	SelfReadState		SelfReadState		`json:"self_read_state"`
	Status			string			`json:"status"`
	NotificationLevel	string			`json:"notification_level"`
	View			[]string		`json:"view"`
	InviterID		ID			`json:"inviter_id"`
	InviteTimestamp		string			`json:"invite_timestamp"`
	SortTimestamp		string			`json:"sort_timestamp"`
	ActiveTimestamp		string			`json:"active_timestamp"`
	DeliveryMediumOption	DeliveryMediumOption	`json:"delivery_medium_option"`
	IsGuest bool `json:"is_guest"`
}

type ParticipantData []struct {
	ID			ParticipantID	`json:"id"`
	FallbackName		string		`json:"fallback_name"`
	InvitationStatus	string		`json:"invitation_status"`
	ParticipantType		string		`json:"participant_type"`
	NewInvitationStatus	string		`json:"new_invitation_status"`
}

type Conversation struct {
	ID                     ID		`json:"id"`
	Type                   string		`json:"type"`
	SelfConversationState			`json:"self_conversation_state"`
	ReadState              []SelfReadState	`json:"read_state"`
	HasActiveHangout       bool		`json:"has_active_hangout"`
	OtrStatus              string		`json:"otr_status"`
	OtrToggle              string		`json:"otr_toggle"`
	CurrentParticipant     []ParticipantID	`json:"current_participant"`
	ParticipantData        ParticipantData	`json:"participant_data"`
	ForkOnExternalInvite   bool		`json:"fork_on_external_invite"`
	NetworkType            []string		`json:"network_type"`
	ForceHistoryState      string		`json:"force_history_state"`
	GroupLinkSharingStatus string		`json:"group_link_sharing_status"`
}

type Formatting struct {
	Bold		bool	`json:"bold"`
	Italics		bool	`json:"italics"`
	Strikethrough	bool	`json:"strikethrough"`
	Underline	bool	`json:"underline"`
}

type Segment []struct {
	Type		string	`json:"type"`
	Text		string	`json:"text"`
	Formatting		`json:"formatting"`
}

type MessageContent struct {
	Segment		Segment	`json:"segment"`
}

type ChatMessage struct {
	MessageContent	MessageContent	`json:"message_content"`
}

type MembershipChange struct {
	Type		string		`json:"type"`
	ParticipantID	[]ParticipantID	`json:"participant_id"`
	LeaveReason	string		`json:"leave_reason"`
}

type SelfEventState struct {
	UserID			ID	`json:"user_id"`
	NotificationLevel	string	`json:"notification_level"`
}

type Event struct {
	ConversationID		ID			`json:"conversation_id"`
	SenderID		ParticipantID		`json:"sender_id"`
	Timestamp		string			`json:"timestamp"`
	SelfEventState		SelfEventState		`json:"self_event_state"`
	ChatMessage		ChatMessage		`json:"chat_message,omitempty"`
	EventID			string			`json:"event_id"`
	AdvancesSortTimestamp	bool			`json:"advances_sort_timestamp"`
	EventOtr		string			`json:"event_otr"`
	DeliveryMedium		DeliveryMedium		`json:"delivery_medium"`
	EventType		string			`json:"event_type"`
	EventVersion		string			`json:"event_version"`
	MembershipChange	MembershipChange	`json:"membership_change,omitempty"`
}

type ConversationState struct {
	ConversationID ID           `json:"conversation_id"`
	Conversation   Conversation `json:"conversation"`
	Event          []Event      `json:"event"`
}

type ConversationHeader struct {
	ConversationID    ID			`json:"conversation_id"`
	ResponseHeader    ResponseHeader	`json:"response_header"`
	ConversationState ConversationState	`json:"conversation_state"`
}

type Hangouts struct {
	ContinuationEndTimestamp	string			`json:"continuation_end_timestamp"`
	ConversationState		[]ConversationHeader	`json:"conversation_state"`
}

