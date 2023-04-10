Các loại docker-compose override cho việc điều chỉnh và mở rộng docker-compose gồm:

* docker-compose.override.yml: Được sử dụng để ghi đè các tùy chọn của docker-compose.yml, ví dụ như thay đổi cổng hoặc biến môi trường.

* docker-compose.<filename>.yml: Được sử dụng để tạo ra các file docker-compose có tính cụ thể hóa cao hơn, ví dụ như tạo một file docker-compose.test.yml để chạy các bài kiểm tra tự động.

* docker-compose.<project-name>.yml: Được sử dụng để thực hiện multiple compose project trên cùng một máy tính, trong đó mỗi project được định nghĩa bởi một file docker-compose có tên riêng biệt.

* docker-compose.<service-name>.yml: Được sử dụng để định nghĩa các tùy chọn riêng cho một service nào đó, như cổng hoặc volume.

Chú ý rằng các file override phải được đặt trong cùng thư mục với file docker-compose.yml, và các tùy chọn của file override sẽ ghi đè lên tùy chọn tương ứng trong file docker-compose.yml.