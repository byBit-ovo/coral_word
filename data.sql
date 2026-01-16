
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



SELECT lr.word_id, lr.familiarity, lr.consecutive_correct, lr.next_review_time FROM learning_record lr JOIN vocabulary v ON lr.word_id = v.id WHERE lr.user_id = '64a3a609-85d3-44ff-8f41-4efcd7a4a975' AND lr.book_id = 'a758ac1b-029a-44f8-a3e8-5e3646a2e6e5' AND lr.next_review_time <= NOW() ORDER BY lr.familiarity ASC, lr.next_review_time ASC LIMIT 10