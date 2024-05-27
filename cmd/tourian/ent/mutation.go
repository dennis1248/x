// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"within.website/x/cmd/tourian/ent/chatmessage"
	"within.website/x/cmd/tourian/ent/conversation"
	"within.website/x/cmd/tourian/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeChatMessage  = "ChatMessage"
	TypeConversation = "Conversation"
)

// ChatMessageMutation represents an operation that mutates the ChatMessage nodes in the graph.
type ChatMessageMutation struct {
	config
	op              Op
	typ             string
	id              *string
	conversation_id *string
	role            *string
	content         *string
	clearedFields   map[string]struct{}
	done            bool
	oldValue        func(context.Context) (*ChatMessage, error)
	predicates      []predicate.ChatMessage
}

var _ ent.Mutation = (*ChatMessageMutation)(nil)

// chatmessageOption allows management of the mutation configuration using functional options.
type chatmessageOption func(*ChatMessageMutation)

// newChatMessageMutation creates new mutation for the ChatMessage entity.
func newChatMessageMutation(c config, op Op, opts ...chatmessageOption) *ChatMessageMutation {
	m := &ChatMessageMutation{
		config:        c,
		op:            op,
		typ:           TypeChatMessage,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withChatMessageID sets the ID field of the mutation.
func withChatMessageID(id string) chatmessageOption {
	return func(m *ChatMessageMutation) {
		var (
			err   error
			once  sync.Once
			value *ChatMessage
		)
		m.oldValue = func(ctx context.Context) (*ChatMessage, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().ChatMessage.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withChatMessage sets the old ChatMessage of the mutation.
func withChatMessage(node *ChatMessage) chatmessageOption {
	return func(m *ChatMessageMutation) {
		m.oldValue = func(context.Context) (*ChatMessage, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ChatMessageMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ChatMessageMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of ChatMessage entities.
func (m *ChatMessageMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ChatMessageMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ChatMessageMutation) IDs(ctx context.Context) ([]string, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []string{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().ChatMessage.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetConversationID sets the "conversation_id" field.
func (m *ChatMessageMutation) SetConversationID(s string) {
	m.conversation_id = &s
}

// ConversationID returns the value of the "conversation_id" field in the mutation.
func (m *ChatMessageMutation) ConversationID() (r string, exists bool) {
	v := m.conversation_id
	if v == nil {
		return
	}
	return *v, true
}

// OldConversationID returns the old "conversation_id" field's value of the ChatMessage entity.
// If the ChatMessage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMessageMutation) OldConversationID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldConversationID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldConversationID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldConversationID: %w", err)
	}
	return oldValue.ConversationID, nil
}

// ResetConversationID resets all changes to the "conversation_id" field.
func (m *ChatMessageMutation) ResetConversationID() {
	m.conversation_id = nil
}

// SetRole sets the "role" field.
func (m *ChatMessageMutation) SetRole(s string) {
	m.role = &s
}

// Role returns the value of the "role" field in the mutation.
func (m *ChatMessageMutation) Role() (r string, exists bool) {
	v := m.role
	if v == nil {
		return
	}
	return *v, true
}

// OldRole returns the old "role" field's value of the ChatMessage entity.
// If the ChatMessage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMessageMutation) OldRole(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRole is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRole requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRole: %w", err)
	}
	return oldValue.Role, nil
}

// ResetRole resets all changes to the "role" field.
func (m *ChatMessageMutation) ResetRole() {
	m.role = nil
}

// SetContent sets the "content" field.
func (m *ChatMessageMutation) SetContent(s string) {
	m.content = &s
}

// Content returns the value of the "content" field in the mutation.
func (m *ChatMessageMutation) Content() (r string, exists bool) {
	v := m.content
	if v == nil {
		return
	}
	return *v, true
}

// OldContent returns the old "content" field's value of the ChatMessage entity.
// If the ChatMessage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMessageMutation) OldContent(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldContent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldContent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContent: %w", err)
	}
	return oldValue.Content, nil
}

// ResetContent resets all changes to the "content" field.
func (m *ChatMessageMutation) ResetContent() {
	m.content = nil
}

// Where appends a list predicates to the ChatMessageMutation builder.
func (m *ChatMessageMutation) Where(ps ...predicate.ChatMessage) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ChatMessageMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ChatMessageMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.ChatMessage, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ChatMessageMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ChatMessageMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (ChatMessage).
func (m *ChatMessageMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ChatMessageMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.conversation_id != nil {
		fields = append(fields, chatmessage.FieldConversationID)
	}
	if m.role != nil {
		fields = append(fields, chatmessage.FieldRole)
	}
	if m.content != nil {
		fields = append(fields, chatmessage.FieldContent)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ChatMessageMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case chatmessage.FieldConversationID:
		return m.ConversationID()
	case chatmessage.FieldRole:
		return m.Role()
	case chatmessage.FieldContent:
		return m.Content()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ChatMessageMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case chatmessage.FieldConversationID:
		return m.OldConversationID(ctx)
	case chatmessage.FieldRole:
		return m.OldRole(ctx)
	case chatmessage.FieldContent:
		return m.OldContent(ctx)
	}
	return nil, fmt.Errorf("unknown ChatMessage field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ChatMessageMutation) SetField(name string, value ent.Value) error {
	switch name {
	case chatmessage.FieldConversationID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetConversationID(v)
		return nil
	case chatmessage.FieldRole:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRole(v)
		return nil
	case chatmessage.FieldContent:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContent(v)
		return nil
	}
	return fmt.Errorf("unknown ChatMessage field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ChatMessageMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ChatMessageMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ChatMessageMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown ChatMessage numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ChatMessageMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ChatMessageMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ChatMessageMutation) ClearField(name string) error {
	return fmt.Errorf("unknown ChatMessage nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ChatMessageMutation) ResetField(name string) error {
	switch name {
	case chatmessage.FieldConversationID:
		m.ResetConversationID()
		return nil
	case chatmessage.FieldRole:
		m.ResetRole()
		return nil
	case chatmessage.FieldContent:
		m.ResetContent()
		return nil
	}
	return fmt.Errorf("unknown ChatMessage field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ChatMessageMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ChatMessageMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ChatMessageMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ChatMessageMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ChatMessageMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ChatMessageMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ChatMessageMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown ChatMessage unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ChatMessageMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown ChatMessage edge %s", name)
}

// ConversationMutation represents an operation that mutates the Conversation nodes in the graph.
type ConversationMutation struct {
	config
	op            Op
	typ           string
	id            *string
	page_url      *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Conversation, error)
	predicates    []predicate.Conversation
}

var _ ent.Mutation = (*ConversationMutation)(nil)

// conversationOption allows management of the mutation configuration using functional options.
type conversationOption func(*ConversationMutation)

// newConversationMutation creates new mutation for the Conversation entity.
func newConversationMutation(c config, op Op, opts ...conversationOption) *ConversationMutation {
	m := &ConversationMutation{
		config:        c,
		op:            op,
		typ:           TypeConversation,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withConversationID sets the ID field of the mutation.
func withConversationID(id string) conversationOption {
	return func(m *ConversationMutation) {
		var (
			err   error
			once  sync.Once
			value *Conversation
		)
		m.oldValue = func(ctx context.Context) (*Conversation, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Conversation.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withConversation sets the old Conversation of the mutation.
func withConversation(node *Conversation) conversationOption {
	return func(m *ConversationMutation) {
		m.oldValue = func(context.Context) (*Conversation, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ConversationMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ConversationMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Conversation entities.
func (m *ConversationMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ConversationMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ConversationMutation) IDs(ctx context.Context) ([]string, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []string{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Conversation.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetPageURL sets the "page_url" field.
func (m *ConversationMutation) SetPageURL(s string) {
	m.page_url = &s
}

// PageURL returns the value of the "page_url" field in the mutation.
func (m *ConversationMutation) PageURL() (r string, exists bool) {
	v := m.page_url
	if v == nil {
		return
	}
	return *v, true
}

// OldPageURL returns the old "page_url" field's value of the Conversation entity.
// If the Conversation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ConversationMutation) OldPageURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPageURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPageURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPageURL: %w", err)
	}
	return oldValue.PageURL, nil
}

// ResetPageURL resets all changes to the "page_url" field.
func (m *ConversationMutation) ResetPageURL() {
	m.page_url = nil
}

// Where appends a list predicates to the ConversationMutation builder.
func (m *ConversationMutation) Where(ps ...predicate.Conversation) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ConversationMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ConversationMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Conversation, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ConversationMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ConversationMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Conversation).
func (m *ConversationMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ConversationMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.page_url != nil {
		fields = append(fields, conversation.FieldPageURL)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ConversationMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case conversation.FieldPageURL:
		return m.PageURL()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ConversationMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case conversation.FieldPageURL:
		return m.OldPageURL(ctx)
	}
	return nil, fmt.Errorf("unknown Conversation field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ConversationMutation) SetField(name string, value ent.Value) error {
	switch name {
	case conversation.FieldPageURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPageURL(v)
		return nil
	}
	return fmt.Errorf("unknown Conversation field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ConversationMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ConversationMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ConversationMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Conversation numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ConversationMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ConversationMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ConversationMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Conversation nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ConversationMutation) ResetField(name string) error {
	switch name {
	case conversation.FieldPageURL:
		m.ResetPageURL()
		return nil
	}
	return fmt.Errorf("unknown Conversation field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ConversationMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ConversationMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ConversationMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ConversationMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ConversationMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ConversationMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ConversationMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Conversation unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ConversationMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Conversation edge %s", name)
}
