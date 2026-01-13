
use coral_word;
CREATE TABLE word_learning (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id VARCHAR(64) NOT NULL,
    word_id BIGINT NOT NULL,
    book_id BIGINT NOT NULL,
    
    familiarity INT DEFAULT 0,       
    consecutive_correct INT DEFAULT 0,   
    total_reviews INT DEFAULT 0,        
    correct_count INT DEFAULT 0,        
    wrong_count INT DEFAULT 0,          
    
  
    last_review_time DATETIME,          
    next_review_time DATETIME,           
    first_learn_time DATETIME,           
    
   
    today_reviews INT DEFAULT 0,        
    today_correct INT DEFAULT 0,         
    
    UNIQUE KEY unique_user_word (user_id, word_id),
    INDEX idx_next_review (user_id, book_id, next_review_time)
);

INSERT INTO word_learning (
    user_id,
    book_id,
    word_id,
    familiarity,
    consecutive_correct,
    total_reviews,
    correct_count,
    wrong_count,
    last_review_time,
    next_review_time,
    today_reviews,
    today_correct
) VALUES (
    ?, ?, ?,          -- user_id, book_id, word_id
    0, 0, 0, 0, 0,    -- 新词初始化
    NULL, NOW(),      -- last_review_time 为空或 NOW()（看你业务），next_review_time 默认 NOW() 代表立即可复习
    0, 0
);