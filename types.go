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

type Thumbnail struct {
	URL      	string		`json:"url"`
	ImageURL 	string		`json:"image_url"`
	WidthPx  	int			`json:"width_px"`
	HeightPx 	int			`json:"height_px"`
}
type EmbedsPlusPhotoPlusPhoto struct {
	Thumbnail 			Thumbnail 	`json:"thumbnail"`
	OwnerObfuscatedID	string   	`json:"owner_obfuscated_id"`
	AlbumID				string   	`json:"album_id"`
	PhotoID				string   	`json:"photo_id"`
	URL					string   	`json:"url"`
	OriginalContentURL	string   	`json:"original_content_url"`
	MediaType			string   	`json:"media_type"`
	StreamID			[]string 	`json:"stream_id"`
} 

type EmbedsPlusAudioV2PlusAudioV2 struct {
	URL					string		`json:"url"`
	OwnerObfuscatedID	string		`json:"owner_obfuscated_id"`
	AlbumID				string   	`json:"album_id"`
	PhotoID				string   	`json:"photo_id"`
	EmbedURL			string		`json:"embed_url"`
	Duration			string		`json:"duration"`
	MediaKey			string		`json:"media_key"`		
}


type Address struct {
	EmbedsPostalAddressV2PostalAddressV2 struct {
		StreetAddress string `json:"street_address"`
	} `json:"embeds.PostalAddressV2.postal_address_v2"`
}

type Geo struct {
	EmbedsGeoCoordinatesV2GeoCoordinatesV2 struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"embeds.GeoCoordinatesV2.geo_coordinates_v2"`
}

type RepresentativeImage struct {
	Type                             []string `json:"type"`
	ID                               string   `json:"id"`
	EmbedsImageObjectV2ImageObjectV2 struct {
		URL string `json:"url"`
	} `json:"embeds.ImageObjectV2.image_object_v2"`
}

type EmbedsPlaceV2PlaceV2 struct {
	URL     				string 					`json:"url"`
	Name    				string 					`json:"name"`
	Address 				Address  				`json:"address"`
	Geo  					Geo						`json:"geo"`
	RepresentativeImage 	RepresentativeImage		`json:"representative_image"`
}

type RepresentativeImage struct {
	Type                             []string `json:"type"`
	ID                               string   `json:"id"`
	EmbedsImageObjectV2ImageObjectV2 struct {
		URL string `json:"url"`
	} `json:"embeds.ImageObjectV2.image_object_v2"`
}

type EmbedsThingV2ThingV2 struct {
	URL                 string `json:"url"`
	Name                string `json:"name"`
	RepresentativeImage RepresentativeImage `json:"representative_image"`
}

type EmbedItem struct {
	Type                     []string `json:"type"`
	ID                       string   `json:"id"`
	EmbedsPlusPhotoPlusPhoto EmbedsPlusPhotoPlusPhoto `json:"embeds.PlusPhoto.plus_photo"`
	EmbedsPlusAudioV2PlusAudioV2 EmbedsPlusAudioV2PlusAudioV2 `json:"embeds.PlusAudioV2.plus_audio_v2"`
	EmbedsPlaceV2PlaceV2  EmbedsPlaceV2PlaceV2 `json:"embeds.PlaceV2.place_v2"`
	EmbedsThingV2ThingV2	EmbedsThingV2ThingV2 `json:"embeds.ThingV2.thing_v2"`
} 

type Attachment []struct {
    EmbedItem 		EmbedItem 		`json:"embed_item"`	
	ID 				string 			`json:"id"`
}

type MessageContent struct {
	Segment			Segment			`json:"segment"`
	Attachment  	Attachment 		`json:"attachment"`
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

