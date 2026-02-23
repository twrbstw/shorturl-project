CREATE TABLE short_urls (
    id SERIAL PRIMARY KEY,
    code VARCHAR(20) UNIQUE NOT NULL,
    long_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expired_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP + INTERVAL '10 minutes'
);

CREATE INDEX idx_short_urls_code ON short_urls(code);
CREATE INDEX idx_short_urls_expired_partial ON short_urls (expires_at) WHERE expires_at IS NOT NULL;

