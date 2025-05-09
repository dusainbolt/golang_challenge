# Các nguyên tắc trong Uber Go Style Guide

Nguồn tài liệu: https://github.com/nc-minh/uber-go-guide-vi

## Pointers to Interfaces

- Không nên dùng con trỏ đến interface vì interface vốn đã chứa con trỏ của chính nó.

## Xác minh sự tuân thủ Interfaces

- Dùng khai báo ẩn danh như `var _ Interface = (*Type)(nil)` để xác minh tại thời điểm biên dịch.

## Receivers and Interfaces

- Chọn value hoặc pointer receiver phù hợp với việc có cần thay đổi dữ liệu hay không.

## Muexes có giá trị bằng 0 là hợp lệ

- Giá trị mặc định của `sync.Mutex` là hợp lệ, không cần dùng con trỏ.

## Sao chép Slices and Maps tại Boundaries

- Luôn sao chép `slice` hoặc `map` khi truyền qua goroutine hoặc giữa các thành phần để tránh lỗi race.

## Dùng Defer để dọn dẹp

- Sử dụng `defer` để đảm bảo giải phóng tài nguyên, ngay cả khi hàm return sớm.

## Kích thước Channel là 1 hoặc None

- Sử dụng channel không buffer hoặc buffer 1 để đơn giản hóa đồng bộ.

## Bắt đầu Enums tại 1

- Tránh dùng giá trị mặc định là 0 cho enum vì 0 thường mang ý nghĩa đặc biệt.

## Sử "time" để xử lý thời gian

- Dùng package `time` để tránh các lỗi không mong muốn khi xử lý thời gian và thời lượng.

## Lỗi

- Lỗi nên được trả về, không panic. Sử dụng `errors.Is`, `errors.As` để xử lý chính xác.

## Các loại lỗi

- Dùng `errors.New`, `fmt.Errorf`, hoặc error type tùy vào nhu cầu matching hoặc ngữ cảnh lỗi.

## Gói lỗi

- Dùng `%w` để gói lỗi giúp giữ nguyên lỗi gốc, hỗ trợ `errors.Is`.

## Lỗi đặt tên

- Biến lỗi xuất khẩu nên bắt đầu bằng `Err`, lỗi không xuất thì dùng `err`.

## Xử lý Lỗi Một Lần

- Không ghi log và return lỗi đồng thời, để tránh trùng lặp xử lý ở tầng cao.

## Xử Lý Lỗi Kiểu Assertion

- Luôn dùng dạng `comma ok` khi ép kiểu để tránh panic.

## Không Sử Dụng Panic

- Chỉ dùng panic khi gặp lỗi không thể phục hồi, không nên panic trong logic chính.

## Sử Dụng go.uber.org/atomic
 
- Dùng `atomic` wrapper của Uber để thao tác thread-safe với boolean, int, ...

## Tránh Sử Dụng Biến Toàn Cục Thay Đổi

- Sử dụng dependency injection thay vì biến toàn cục có thể thay đổi.

## Tránh Nhúng Các Loại vào Cấu Trúc Công Khai

- Không nhúng struct khác vào public struct vì sẽ rò rỉ chi tiết triển khai.

## Tránh Sử Dụng Tên Được Xây Dựng Sẵn

- Không đặt tên biến là `string`, `error`, `int`... để tránh shadow tên gốc.

## Tránh init()

- Không nên dùng `init()` vì dễ gây lỗi khó debug xung đột với base, nên dùng hàm khởi tạo rõ ràng prefix => (default, load....)

## Thoát Ở Main

- Chỉ dùng `os.Exit` hoặc `log.Fatal` trong `main()` để dễ test và clean-up.

## Thoát Một Lần

- Toàn bộ việc thoát chương trình nên tập trung ở một chỗ duy nhất trong `main()`.

## Sử Dụng Thẻ Trường Trong Các Cấu Trúc Được Unmarshaled

- Gắn thẻ JSON, YAML cho các trường để dễ serialize và đổi tên an toàn.

## Không Gửi và Quên Goroutines

- Luôn quản lý vòng đời goroutine, tránh tạo mà không có cách kiểm soát.

## Chờ Goroutines Thoát

- Dùng `sync.WaitGroup` hoặc `chan` để chờ goroutines hoàn tất.

## Không Có Goroutines Trong init()

- Không nên khởi tạo goroutine trong `init()`, hãy dùng object quản lý lifecycle.

## Hiệu Suất

- Ưu tiên tối ưu hóa thực sự cần thiết, có đo lường.

## Ưu Tiên strconv hơn fmt

- Dùng `strconv` thay vì `fmt` để chuyển đổi dữ liệu nhanh và ít cấp phát hơn.

## Tránh Chuyển Đổi Chuỗi sang Byte Lặp Lại

- Tránh tạo `[]byte` từ cùng một chuỗi nhiều lần.

## Ưu Tiên Xác Định Khả Năng của Container

- Chỉ định dung lượng ban đầu cho map và slice để tránh realloc.

## Kiểu

- Giữ mã ngắn gọn, nhất quán, và dễ đọc.

## Tránh Các Dòng Quá Dài

- Giới hạn dòng dưới 100 ký tự để dễ đọc.

## Duy Trì Sự Nhất Quán

- Tuân theo phong cách đã chọn trong toàn bộ codebase.

## Nhóm Các Khai Báo Tương Tự

- Khai báo biến/hàm có mục đích giống nhau nên đặt gần nhau.

## Thứ Tự Nhóm Nhập

- Nhập theo nhóm: chuẩn, ngoài, nội bộ.

## Tên Gói

- Tên gói nên ngắn, viết thường, không có dấu gạch dưới.

## Tên Hàm

- Tên hàm nên rõ ràng, mô tả đúng hành vi.

## Bí Danh Nhập

- Chỉ đặt alias khi có xung đột hoặc tên không rõ ràng.

## Nhóm và Sắp Xếp Hàm

- Sắp xếp hàm theo nhóm logic hoặc mức độ truy cập.

## Giảm Lồng Ghép

- Tách logic ra hàm riêng để giảm lồng nhau.

## Else Không Cần Thiết

- Tránh dùng `else` nếu `if` đã `return`.

## Khai Báo Biến Cấp Cao

- Biến cấp cao nên khai báo rõ ràng, tránh dùng chung không rõ ràng.

## Thêm Tiền Tố Cho Các Biến Toàn Cục Không Được Xuất _

- Dùng `_` để phân biệt biến toàn cục không export.

## Nhúng Trong Cấu Trúc

- Chỉ nhúng nếu thực sự cần kế thừa hành vi.

## Khai Báo Biến Cục Bộ

- Biến cục bộ nên được khai báo gần nơi sử dụng.

## nil Là Một Slice Hợp Lệ

- Không cần khởi tạo slice rỗng bằng `[]T{}` nếu có thể dùng `nil`.

## Giảm Phạm Vi của Biến

- Khai báo biến gần nơi sử dụng để hạn chế phạm vi.

## Tránh Tham Số Trần

- Dùng biến đặt tên rõ ràng thay vì truyền số hoặc chuỗi trần.

## Sử Dụng Chuỗi Văn Bản Thô Để Tránh Escape

- Dùng `raw string` (`` `...` ``) để viết chuỗi nhiều dòng dễ hơn.

## Khởi Tạo Các Cấu Trúc

- Dùng tên trường khi khởi tạo struct để rõ ràng.

## Sử Dụng Tên Trường để Khởi Tạo Cấu Trúc

- Giúp đọc và bảo trì code dễ hơn.

## Bỏ Qua Các Trường Giá Trị Zero trong Cấu Trúc

- Không cần gán lại giá trị mặc định (zero value).

## Sử Dụng var Cho Cấu Trúc Giá Trị Zero

- Dùng `var x T` thay vì `x := T{}` để rõ là zero value.

## Khởi Tạo Các Tham Chiếu Cấu Trúc

- Dùng con trỏ khi cần thay đổi giá trị struct.

## Khởi tạo Maps

- Dùng `make(map[K]V, size)` khi biết trước số lượng phần tử.

## Định Dạng Chuỗi Ngoài Printf

- Tách logic format chuỗi ra khỏi `Printf` để dễ kiểm thử.

## Đặt Tên Các Hàm Theo Phong Cách Printf

- Hàm nhận format string nên có hậu tố `f` như `Logf`, `Errorf`.

## Mẫu Thiết Kế

- Sử dụng các mẫu quen thuộc như functional options hoặc test table.

## Bảng Kiểm Tra

- Viết test dạng bảng để kiểm tra nhiều trường hợp logic.

## Tuỳ Chọn Chức Năng

- Dùng functional options để cấu hình đối tượng linh hoạt.

## Kiểm Tra Cú Pháp

- Luôn chạy `go vet`, `golint`, và thiết lập IDE kiểm tra tự động.
