// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: jsonfeed.proto

package jsonfeed

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Feed is the root of a Proto Feed document. A feed must at least contain a
// title and items.
type Feed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (required, string) is the name of the feed, which will often correspond to
	// the name of the website (blog, for instance), though not necessarily.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// (optional but strongly recommended, string) is the URL of the resource that
	// the feed describes. This resource may or may not actually be a “home” page,
	// but it should be an HTML page. If a feed is published on the public web,
	// this should be considered as required. But it may not make sense in the
	// case of a file created on a desktop computer, when that file is not shared
	// or is shared only privately.
	HomePageUrl string `protobuf:"bytes,3,opt,name=home_page_url,json=homePageUrl,proto3" json:"home_page_url,omitempty"`
	// (optional but strongly recommended, string) is the URL of the feed, and
	// serves as the unique identifier for the feed. As with home_page_url, this
	// should be considered required for feeds on the public web.
	FeedUrl string `protobuf:"bytes,4,opt,name=feed_url,json=feedUrl,proto3" json:"feed_url,omitempty"`
	// (optional, string) provides more detail, beyond the title, on what the feed
	// is about. A feed reader may display this text.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// (optional, string) is a description of the purpose of the feed. This is for
	// the use of people looking at the raw Protobuf, and should be ignored by
	// feed readers.
	UserComment string `protobuf:"bytes,6,opt,name=user_comment,json=userComment,proto3" json:"user_comment,omitempty"`
	// (optional, string) is the URL of an image for the feed suitable to be used
	// in a source list. It should be square and relatively large — such as 512 x
	// 512 — so that it can be scaled down and so that it can look good on retina
	// displays. It should use transparency where appropriate, since it may be
	// rendered on a non-white background.
	Icon string `protobuf:"bytes,7,opt,name=icon,proto3" json:"icon,omitempty"`
	// (optional, string) is the URL of an image for the feed suitable to be used
	// in a source list. It should be square and relatively small, but not smaller
	// than 64 x 64.
	Favicon string `protobuf:"bytes,8,opt,name=favicon,proto3" json:"favicon,omitempty"`
	// (optional, array of objects) specifies the feed authors.
	Authors []*Author `protobuf:"bytes,9,rep,name=authors,proto3" json:"authors,omitempty"`
	// (optional, string) is the primary language for the feed.
	Language string `protobuf:"bytes,10,opt,name=language,proto3" json:"language,omitempty"`
	// (optional, boolean) says whether or not the feed is finished — that is,
	// whether or not it will ever update again. A feed for a temporary event,
	// such as an instance of the Olympics, could expire. If the value is true,
	// then it’s expired. Any other value, or the absence of expired, means the
	// feed may continue to update.
	Expired bool `protobuf:"varint,11,opt,name=expired,proto3" json:"expired,omitempty"`
	// (required, array of objects) contains the items in the feed. This is the
	// most important element of the feed after the version field. Each item is a
	// story, blog post, article, photograph, video, or other thing. For example,
	// if a feed contains a long article, a podcast episode, and a photo, those
	// three items would be included in items.
	Items []*Item `protobuf:"bytes,12,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Feed) Reset() {
	*x = Feed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsonfeed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feed) ProtoMessage() {}

func (x *Feed) ProtoReflect() protoreflect.Message {
	mi := &file_jsonfeed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feed.ProtoReflect.Descriptor instead.
func (*Feed) Descriptor() ([]byte, []int) {
	return file_jsonfeed_proto_rawDescGZIP(), []int{0}
}

func (x *Feed) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Feed) GetHomePageUrl() string {
	if x != nil {
		return x.HomePageUrl
	}
	return ""
}

func (x *Feed) GetFeedUrl() string {
	if x != nil {
		return x.FeedUrl
	}
	return ""
}

func (x *Feed) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Feed) GetUserComment() string {
	if x != nil {
		return x.UserComment
	}
	return ""
}

func (x *Feed) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Feed) GetFavicon() string {
	if x != nil {
		return x.Favicon
	}
	return ""
}

func (x *Feed) GetAuthors() []*Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *Feed) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Feed) GetExpired() bool {
	if x != nil {
		return x.Expired
	}
	return false
}

func (x *Feed) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

// Author is an object representing the author of the feed or item.
type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (optional, string) is the author’s name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (optional, string) is the URL of a site owned by the author. It could be a
	// blog, micro-blog, Twitter account, and so on. Ideally the linked-to page
	// provides a way to contact the author, but that’s not required. The URL
	// could be a mailto: link, though we suspect that will be rare.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// (optional, string) is the URL for an image for the author. As with icon, it
	// should be square and relatively large — such as 512 x 512 pixels — and
	// should use transparency where appropriate, since it may be rendered on a
	// non-white background.
	Avatar string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsonfeed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_jsonfeed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_jsonfeed_proto_rawDescGZIP(), []int{1}
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Author) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

// Item is an object representing a single story, blog post, article,
// photograph, video, or other thing within a feed.
type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (required, string) is unique for that item for that feed over time. If an
	// item is ever updated, the id should be unchanged. New items should never
	// use a previously-used id. If an id is presented as a number or other type,
	// a JSON Feed reader must coerce it to a string. Ideally, the id is the full
	// URL of the resource described by the item, since URLs make great unique
	// identifiers.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// (optional, string) is the URL of the resource described by the item. It’s
	// the permalink. This may be the same as the id — but should be present
	// regardless.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// (optional, string) is the URL of a page elsewhere. This is especially
	// useful for linkblogs. If url links to where you’re talking about a thing,
	// then external_url links to the thing you’re talking about.
	ExternalUrl string `protobuf:"bytes,3,opt,name=external_url,json=externalUrl,proto3" json:"external_url,omitempty"`
	// (optional, string) is plain text. Microblog items in particular may omit
	// titles.
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	// (optional, string) is the body of the item. It can be plain text, HTML, or
	// a snippet of Markdown. (It should not be the entire Markdown document; just
	// a snippet.) This is complete enough that it can be displayed alone in a
	// reader.
	ContentText string `protobuf:"bytes,5,opt,name=content_text,json=contentText,proto3" json:"content_text,omitempty"`
	// (optional, string) is the body of the item. It can be plain text, HTML, or
	// a snippet of Markdown. (It should not be the entire Markdown document; just
	// a snippet.) This is complete enough that it can be displayed alone in a
	// reader.
	ContentHtml string `protobuf:"bytes,6,opt,name=content_html,json=contentHtml,proto3" json:"content_html,omitempty"`
	// (optional, string) is a plain text sentence or two describing the item.
	// This might be presented in a timeline, for instance, where a detail view
	// would display all of content_html or content_text.
	Summary string `protobuf:"bytes,7,opt,name=summary,proto3" json:"summary,omitempty"`
	// (optional, string) is the URL of the main image for the item. This image
	// may also appear in the content_html — if so, it’s a hint to the feed reader
	// that this is the main, featured image. Even if it’s not, it will appear in
	// the detail view. Images should be square, with a 4:3 aspect ratio. (We will
	// be flexible on this in the future.)
	Image string `protobuf:"bytes,8,opt,name=image,proto3" json:"image,omitempty"`
	// (optional, string) is the URL of an image to use as a banner. Some blogging
	// systems (such as Medium) display a different banner image in the list view
	// from the detail view. In those systems, this image should be used in the
	// list view, and image in the detail view.
	BannerImage string `protobuf:"bytes,9,opt,name=banner_image,json=bannerImage,proto3" json:"banner_image,omitempty"`
	// (optional, string) specifies the date in RFC 3339 format.
	DatePublished *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=date_published,json=datePublished,proto3" json:"date_published,omitempty"`
	// (optional, string) specifies the modification date in RFC 3339 format.
	DateModified *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=date_modified,json=dateModified,proto3" json:"date_modified,omitempty"`
	// (optional, array of objects) has the same structure as the top-level
	// authors. If not specified in an item, then the top-level authors, if
	// present, are the authors of the item.
	Authors []*Author `protobuf:"bytes,12,rep,name=authors,proto3" json:"authors,omitempty"`
	// (optional, array of strings) can have any plain text values you want. Tags
	// tend to be just one word, but they may be anything. Note: they are not the
	// equivalent of Twitter hashtags. Some blogging systems and other feed
	// formats call these categories.
	Tags []string `protobuf:"bytes,13,rep,name=tags,proto3" json:"tags,omitempty"`
	// (optional, string) is the language for this item, using the same format as
	// the top-level language field. The value can be different than the primary
	// language for the feed when a specific item is written in a different
	// language than other items in the feed.
	Language string `protobuf:"bytes,14,opt,name=language,proto3" json:"language,omitempty"`
	// (optional, array of objects) specifies the attachments associated with the
	// item. Attachments are files that are associated with an item. The value of
	// the attachments field is an array of objects, each of which has a url
	// field, and other fields as specified in the attachment object definition.
	Attachments []*Attachment `protobuf:"bytes,15,rep,name=attachments,proto3" json:"attachments,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsonfeed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_jsonfeed_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_jsonfeed_proto_rawDescGZIP(), []int{2}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Item) GetExternalUrl() string {
	if x != nil {
		return x.ExternalUrl
	}
	return ""
}

func (x *Item) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Item) GetContentText() string {
	if x != nil {
		return x.ContentText
	}
	return ""
}

func (x *Item) GetContentHtml() string {
	if x != nil {
		return x.ContentHtml
	}
	return ""
}

func (x *Item) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Item) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Item) GetBannerImage() string {
	if x != nil {
		return x.BannerImage
	}
	return ""
}

func (x *Item) GetDatePublished() *timestamppb.Timestamp {
	if x != nil {
		return x.DatePublished
	}
	return nil
}

func (x *Item) GetDateModified() *timestamppb.Timestamp {
	if x != nil {
		return x.DateModified
	}
	return nil
}

func (x *Item) GetAuthors() []*Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *Item) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Item) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Item) GetAttachments() []*Attachment {
	if x != nil {
		return x.Attachments
	}
	return nil
}

// Attachment is an object representing a file associated with an item.
type Attachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (required, string) specifies the location of the attachment.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// (required, string) specifies the type of the attachment, such as
	// “audio/mpeg.”
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// (optional, string) specifies the title of the attachment.
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// (optional, number) specifies how large the file is.
	SizeInBytes int32 `protobuf:"varint,4,opt,name=size_in_bytes,json=sizeInBytes,proto3" json:"size_in_bytes,omitempty"`
	// (optional, number) specifies how long it takes to listen to or watch, when
	// played at normal speed.
	DurationInSeconds int32 `protobuf:"varint,5,opt,name=duration_in_seconds,json=durationInSeconds,proto3" json:"duration_in_seconds,omitempty"`
}

func (x *Attachment) Reset() {
	*x = Attachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsonfeed_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attachment) ProtoMessage() {}

func (x *Attachment) ProtoReflect() protoreflect.Message {
	mi := &file_jsonfeed_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attachment.ProtoReflect.Descriptor instead.
func (*Attachment) Descriptor() ([]byte, []int) {
	return file_jsonfeed_proto_rawDescGZIP(), []int{3}
}

func (x *Attachment) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Attachment) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *Attachment) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Attachment) GetSizeInBytes() int32 {
	if x != nil {
		return x.SizeInBytes
	}
	return 0
}

func (x *Attachment) GetDurationInSeconds() int32 {
	if x != nil {
		return x.DurationInSeconds
	}
	return 0
}

var File_jsonfeed_proto protoreflect.FileDescriptor

var file_jsonfeed_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6a, 0x73, 0x6f, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x02, 0x0a, 0x04,
	0x46, 0x65, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x68, 0x6f,
	0x6d, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x68, 0x6f, 0x6d, 0x65, 0x50, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x19,
	0x0a, 0x08, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x66, 0x65, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x76, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x76, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x07,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x6a, 0x73, 0x6f, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52,
	0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x12, 0x24,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x6a, 0x73, 0x6f, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x46, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x92, 0x04, 0x0a,
	0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54,
	0x65, 0x78, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x68,
	0x74, 0x6d, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x48, 0x74, 0x6d, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x0d,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0c, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x2a, 0x0a,
	0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0xa5, 0x01, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x69, 0x6e,
	0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x69,
	0x7a, 0x65, 0x49, 0x6e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x42, 0x2a, 0x5a, 0x28, 0x77, 0x69, 0x74,
	0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2f, 0x78, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jsonfeed_proto_rawDescOnce sync.Once
	file_jsonfeed_proto_rawDescData = file_jsonfeed_proto_rawDesc
)

func file_jsonfeed_proto_rawDescGZIP() []byte {
	file_jsonfeed_proto_rawDescOnce.Do(func() {
		file_jsonfeed_proto_rawDescData = protoimpl.X.CompressGZIP(file_jsonfeed_proto_rawDescData)
	})
	return file_jsonfeed_proto_rawDescData
}

var file_jsonfeed_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_jsonfeed_proto_goTypes = []any{
	(*Feed)(nil),                  // 0: jsonfeed.Feed
	(*Author)(nil),                // 1: jsonfeed.Author
	(*Item)(nil),                  // 2: jsonfeed.Item
	(*Attachment)(nil),            // 3: jsonfeed.Attachment
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_jsonfeed_proto_depIdxs = []int32{
	1, // 0: jsonfeed.Feed.authors:type_name -> jsonfeed.Author
	2, // 1: jsonfeed.Feed.items:type_name -> jsonfeed.Item
	4, // 2: jsonfeed.Item.date_published:type_name -> google.protobuf.Timestamp
	4, // 3: jsonfeed.Item.date_modified:type_name -> google.protobuf.Timestamp
	1, // 4: jsonfeed.Item.authors:type_name -> jsonfeed.Author
	3, // 5: jsonfeed.Item.attachments:type_name -> jsonfeed.Attachment
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_jsonfeed_proto_init() }
func file_jsonfeed_proto_init() {
	if File_jsonfeed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jsonfeed_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Feed); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_jsonfeed_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Author); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_jsonfeed_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Item); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_jsonfeed_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Attachment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_jsonfeed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_jsonfeed_proto_goTypes,
		DependencyIndexes: file_jsonfeed_proto_depIdxs,
		MessageInfos:      file_jsonfeed_proto_msgTypes,
	}.Build()
	File_jsonfeed_proto = out.File
	file_jsonfeed_proto_rawDesc = nil
	file_jsonfeed_proto_goTypes = nil
	file_jsonfeed_proto_depIdxs = nil
}
