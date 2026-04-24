-- =========================
-- MAILBOXES
-- =========================
CREATE TABLE IF NOT EXISTS mailboxes (
    id TEXT PRIMARY KEY,
    address TEXT NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    -- retention settings
    retention_mode TEXT DEFAULT 'auto',
    expires_after_seconds INTEGER
);

-- =========================
-- MESSAGES
-- =========================
CREATE TABLE IF NOT EXISTS messages (
    id TEXT PRIMARY KEY,
    mailbox_id TEXT NOT NULL,

    from_address TEXT,
    subject TEXT,
    body_text TEXT,
    body_html TEXT,

    raw_path TEXT,

    received_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    -- retention tracking
    viewed_at DATETIME,
    delete_at DATETIME,

    size_bytes INTEGER,

    FOREIGN KEY (mailbox_id) REFERENCES mailboxes(id)
);

-- =========================
-- ATTACHMENTS
-- =========================
CREATE TABLE IF NOT EXISTS attachments (
    id TEXT PRIMARY KEY,
    message_id TEXT NOT NULL,

    filename TEXT,
    content_type TEXT,
    size_bytes INTEGER,

    storage_path TEXT NOT NULL,
    sha256 TEXT,

    FOREIGN KEY (message_id) REFERENCES messages(id)
);

CREATE INDEX IF NOT EXISTS idx_messages_mailbox_id ON messages(mailbox_id);
CREATE INDEX IF NOT EXISTS idx_messages_received_at ON messages(received_at);
CREATE INDEX IF NOT EXISTS idx_attachments_message_id ON attachments(message_id);