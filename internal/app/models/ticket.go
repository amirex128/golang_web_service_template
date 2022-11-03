package models

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// تیکتی که پرنت ایدی ان صفر باشد به عنوان تیکت اصلی نمایش داده میشود و بعد از باز کردن آن تمامی تیکت های که پرنت ایدی ان را داشته باشند بر اساس تاریخ مرتب میشوند
type Ticket struct {
	ID          uint64  `gorm:"primary_key;auto_increment" json:"id"`
	ParentID    uint64  `gorm:"default:0" json:"parent_id"`
	IsAnswer    bool    `gorm:"default:false" json:"is_answer"`
	UserID      uint64  `gorm:"default:0" json:"user_id"`
	User        User    `json:"user"`
	GuestName   string  `json:"guest_name"`
	GuestMobile string  `json:"guest_mobile"`
	Title       string  `json:"title"`
	Body        string  `json:"body"`
	GalleryID   uint64  `json:"gallery_id"`
	Gallery     Gallery `json:"gallery"`
	CreatedAt   string  `json:"created_at"`
}

func (m *MysqlManager) CreateTicket(c *gin.Context, dto DTOs.CreateTicket, userID uint64) error {
	var parentTicket Ticket
	err := m.GetConn().Update("is_answer", true).Where("id = ?", dto.ParentID).First(&parentTicket).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در دریافت تیکت ها",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	if parentTicket.UserID != userID && IsAdmin(c) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در دریافت تیکت ها",
			"error":   "شما اجازه ارسال پاسخ به این تیکت را ندارید",
			"type":    "model",
		})
		return err
	}

	ticket := Ticket{
		UserID:      userID,
		ParentID:    dto.ParentID,
		IsAnswer:    dto.IsAnswer,
		GuestName:   dto.GuestName,
		GuestMobile: dto.GuestMobile,
		Title:       dto.Title,
		Body:        dto.Body,
		GalleryID:   dto.GalleryID,
		CreatedAt:   utils.NowTime(),
	}
	err = m.GetConn().Create(&ticket).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در ثبت تیکت",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return nil
}

func (m *MysqlManager) GetAllTicketWithPagination(c *gin.Context, dto DTOs.IndexTicket, userID uint64) (*DTOs.Pagination, error) {
	conn := m.GetConn()
	var tickets []Ticket
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("tickets", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("title LIKE ? OR body LIKE ? ", "%"+dto.Search+"%", "%"+dto.Search+"%").Order("id DESC")
	}
	err := conn.Where("user_id = ?", userID).Preload("Gallery").Preload("User").Find(&tickets).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در دریافت تیکت ها",
			"error":   err.Error(),
			"type":    "model",
		})
		return pagination, err
	}
	pagination.Data = tickets
	return pagination, nil
}

func (m *MysqlManager) GetTicketWithChildren(c *gin.Context, ticketID uint64) ([]Ticket, error) {
	conn := m.GetConn()
	var tickets []Ticket
	var mainTicket Ticket
	err := conn.Where("id = ? ", ticketID).First(&mainTicket).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در دریافت تیکت ها",
			"error":   err.Error(),
			"type":    "model",
		})
		return nil, err
	}

	err = conn.Where("parent_id = ?", ticketID).Order("created_at").Preload("Gallery").Preload("User").Find(&tickets).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در دریافت تیکت ها",
			"error":   err.Error(),
			"type":    "model",
		})
		return nil, err
	}
	tickets = append([]Ticket{mainTicket}, tickets...)
	return tickets, nil
}
