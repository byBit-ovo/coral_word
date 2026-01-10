CREATE TABLE word_learning (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id VARCHAR(64) NOT NULL,
    word_id BIGINT NOT NULL,
    book_id BIGINT NOT NULL,
    
    -- 学习状态
    familiarity INT DEFAULT 0,           -- 熟悉度 (0-5)
    consecutive_correct INT DEFAULT 0,   -- 连续正确次数
    total_reviews INT DEFAULT 0,         -- 总复习次数
    correct_count INT DEFAULT 0,         -- 正确次数
    wrong_count INT DEFAULT 0,           -- 错误次数
    
    -- 时间相关
    last_review_time DATETIME,           -- 上次复习时间
    next_review_time DATETIME,           -- 下次应复习时间
    first_learn_time DATETIME,           -- 首次学习时间
    
    -- 当日学习状态
    today_reviews INT DEFAULT 0,         -- 今日复习次数
    today_correct INT DEFAULT 0,         -- 今日正确次数
    
    UNIQUE KEY unique_user_word (user_id, word_id),
    INDEX idx_next_review (user_id, book_id, next_review_time)
);