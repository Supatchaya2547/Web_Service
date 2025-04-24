CREATE TABLE courses (
    course_id SERIAL PRIMARY KEY,
    course_name VARCHAR(255) NOT NULL,
    course_desc TEXT,
    thumbnail_url VARCHAR(255), 
    course_type VARCHAR(255), 
    course_instructor VARCHAR(255) NOT NULL, 
    profile_url VARCHAR(255), 
    course_price DECIMAL(10,2) NOT NULL, 
    duration VARCHAR(255), 
    rating DECIMAL(2,1), 
    num_reviews INT, 
    enrollment_count INT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE courses_url(
    course_id INT PRIMARY KEY,
    detail_url VARCHAR(255) UNIQUE,
    FOREIGN KEY (course_id) REFERENCES courses(course_id)
);


CREATE TABLE affiliates (
    affiliate_id SERIAL PRIMARY KEY,
    affiliate_name VARCHAR(255),
    affiliate_email VARCHAR(255) UNIQUE,
    affiliate_password VARCHAR(255)
);

CREATE TABLE affiliate_url (
    url_id SERIAL PRIMARY KEY,
    affiliate_id INT,
    aff_url VARCHAR(255),
    clicks INTEGER DEFAULT 0,
    parameter TEXT,
    FOREIGN KEY (affiliate_id) REFERENCES affiliates(affiliate_id)
);

CREATE TABLE request_logs (
    id SERIAL PRIMARY KEY,
    affiliate_id INT,  -- แก้เป็น INT เพื่อให้ตรงกับตาราง affiliates
    action VARCHAR(255),
    parameter TEXT,
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Click_logs (
    id SERIAL PRIMARY KEY,
    affiliate_id INT NOT NULL,  -- เพิ่ม NOT NULL เพื่อให้แน่ใจว่ามีค่า
    course_id INT NOT NULL,     -- เพิ่ม NOT NULL เพื่อให้แน่ใจว่ามีค่า
    click_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (affiliate_id) REFERENCES affiliates(affiliate_id),
    FOREIGN KEY (course_id) REFERENCES courses(course_id)
);

-- สร้าง function สำหรับอัพเดท created_at โดยอัตโนมัติ
CREATE OR REPLACE FUNCTION created_modified_column()
RETURNS TRIGGER AS $$ 
BEGIN 
    NEW.created_at = now();
    RETURN NEW; 
END;
$$ LANGUAGE plpgsql;

-- สร้าง function สำหรับอัพเดท updated_at โดยอัตโนมัติ
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$ 
BEGIN 
    NEW.updated_at = now(); 
    RETURN NEW; 
END;
$$ LANGUAGE plpgsql;

-- สร้าง trigger สำหรับอัพเดท updated_at
CREATE TRIGGER update_courses_modtime
    BEFORE UPDATE ON courses
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- สร้าง trigger สำหรับบันทึกเวลาเมื่อ insert
CREATE TRIGGER insert_courses_createdtime
    BEFORE INSERT ON courses
    FOR EACH ROW
    EXECUTE FUNCTION created_modified_column();

-- เพิ่มค่า clicks 
CREATE OR REPLACE FUNCTION increment_clicks()
RETURNS TRIGGER AS $$ 
BEGIN
    -- เพิ่มค่าของ clicks ที่บันทึกในตาราง Click_logs
    UPDATE Click_logs 
    SET clicks = clicks + 1
    WHERE affiliate_id = NEW.affiliate_id 
    AND course_id = NEW.course_id;

    -- ถ้าไม่มีแถวที่ตรงกันให้เพิ่มแถวใหม่
    INSERT INTO Click_logs (affiliate_id, course_id, clicks)
    SELECT NEW.affiliate_id, NEW.course_id, 1
    WHERE NOT EXISTS (SELECT 1 FROM Click_logs WHERE affiliate_id = NEW.affiliate_id AND course_id = NEW.course_id);

    -- ส่งคืนค่า NEW เพื่อให้ insert ข้อมูลตามปกติ
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_clicks
AFTER INSERT ON Click_logs
FOR EACH ROW
EXECUTE FUNCTION increment_clicks();


-- เพิ่มข้อมูลตัวอย่าง
INSERT INTO courses (
    course_name, -- ชื่อคอร์ส
    course_desc, -- คำอธิบายเกี่ยวกับคอร์ส
    thumbnail_url, -- URL ของภาพปกหรือรูปตัวอย่างของคอร์ส
    course_type, --ประเภทของคอร์ส 
    course_instructor, --ชื่อของผู้สอน
    profile_url, -- ลิงก์ไปยังโปรไฟล์ของผู้สอน
    course_price, -- ราคาของคอร์ส
    duration, -- ระยะเวลาเรียน
    rating, -- คะแนนรีวิวเฉลี่ย
    num_reviews, -- จำนวนรีวิว
    enrollment_count -- จำนวนผู้ลงทะเบียนเรียนแล้ว
) VALUES 
-- 1
(
    'Python for Beginners',
    'เรียนรู้พื้นฐานการเขียนโปรแกรมด้วยภาษา Python เหมาะสำหรับผู้เริ่มต้น',
    'https://example.com/images/python.jpg',
    'Python',
    'Jane Smith',
    'https://example.com/instructors/jane',
    0.00,
    '6 ชั่วโมง',
    4.8,
    1523,
    28000
),
-- 2
(
    'Excel ขั้นเทพสำหรับงานออฟฟิศ', 
    'ใช้งาน Excel ตั้งแต่พื้นฐานถึงขั้นสูง พร้อมสูตรและเทคนิค', 
    'https://example.com/images/excel.jpg',
    'Excel',
    'ภัทรพล โพธิ์ศรี', 
    'https://example.com/instructors/pat', 
    699.00, 
    '8 ชั่วโมง', 
    4.6, 
    875, 
    12500
),
-- 3
(
    'วาดภาพ Digital Art ด้วย Procreate',
    'เรียนรู้การวาดภาพบน iPad ด้วยแอป Procreate ตั้งแต่พื้นฐาน',
    'https://example.com/images/procreate.jpg',
    'Art',
    'Arisa Chen',
    'https://example.com/instructors/arisa',
    1200.00,
    '5 ชั่วโมง 30 นาที',
    4.9,
    412,
    5300
),
-- 4
(
    'JavaScript Web Development',
    'เขียนเว็บไซต์แบบ Interactive ด้วย JavaScript',
    'https://example.com/images/js.jpg',
    'JavaScript',
    'John Doe',
    'https://example.com/instructors/john',
    0.00,
    '7 ชั่วโมง',
    4.7,
    1399,
     24000
),
-- 5
(
    'การจัดการเวลาอย่างมีประสิทธิภาพ',
    'เรียนรู้เทคนิค Time Management และ Productivity',
    'https://example.com/images/time.jpg',
    'Management',
    'นภัสสร ลิ้มวัฒนา',
    'https://example.com/instructors/napas0',
    499.00,
    '3 ชั่วโมง',
    4.5,
    326,
    8000
),
-- 6
(
    'UI/UX Design สำหรับมือใหม่',
    'เรียนรู้การออกแบบประสบการณ์ผู้ใช้ในแอปและเว็บไซต์',
    'https://example.com/images/uiux.jpg',
    'UX/UI',
    'Kelvin Wong',
    'https://example.com/instructors/kelvin',
    990.00,
    '10 ชั่วโมง',
    4.8,
    978,
    10800
),
-- 7
(
    'ภาษาอังกฤษสำหรับการทำงาน',
    'พัฒนาทักษะภาษาอังกฤษในที่ทำงาน เน้นการสนทนา',
    'https://example.com/images/english.jpg',
    'English',
    'ครูแพร',
    'https://example.com/instructors/kru-prae',
    499.00,
    '4 ชั่วโมง',
    4.4,
    610,
    19400
),
-- 8
(
    'เรียนถ่ายภาพด้วยกล้องมือถือ',
    'ใช้สมาร์ตโฟนถ่ายภาพให้ดูมืออาชีพ',
    'https://example.com/images/photo.jpg',
    'Photography',
    'วรากร สุทธิกุล',
    'https://example.com/instructors/warakorn',
    450.00,
    '3 ชั่วโมง 45 นาที',
    4.6,
    322,
    6700
),
-- 9
(
    'การตลาดออนไลน์ 101',
    'พื้นฐานการทำตลาดบน Facebook, Google และ IG',
    'https://example.com/images/marketing.jpg',
    'Marketing',
    'ศิรดา แสงทอง',
    'https://example.com/instructors/sirada',
    899.00,
    '6 ชั่วโมง',
    4.7,
    801,
    9400
),
-- 10
(
    'เรียนทำอาหารไทยแบบต้นตำรับ',
    'สอนทำเมนูไทยยอดนิยม เช่น ต้มยำ ผัดไทย แกงเขียวหวาน',
    'https://example.com/images/thai-food.jpg',
    'Cooking',
    'เชฟจุฬา',
    'https://example.com/instructors/chef-jula',
    1500.00,
    '6 ชั่วโมง',
    4.9,
    550,
    3900
);

INSERT INTO courses_url(
    course_id,
    detail_url
) VALUES
(1, 'https://example.com/python-beginners'),
(2, 'https://example.com/excel-master'),
(3, 'https://example.com/procreate-course'),
(4, 'https://example.com/js-webdev'),
(5, 'https://example.com/time-management'),
(6, 'https://example.com/uiux-design'),
(7, 'https://example.com/business-english'),
(8, 'https://example.com/mobile-photo'),
(9, 'https://example.com/online-marketing'),
(10, 'https://example.com/thai-cooking');
