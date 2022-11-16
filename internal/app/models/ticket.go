package models

import (
	"github.com/amirex128/selloora_backend/internal/app/DTOs"
	"github.com/amirex128/selloora_backend/internal/app/utils"
	"context"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// تیکتی که پرنت ایدی ان صفر باشد به عنوان تیکت اصلی نمایش داده میشود و بعد از باز کردن آن تمامی تیکت های که پرنت ایدی ان را داشته باشند بر اساس تاریخ مرتب میشوند
type Ticket struct {
	ID          uint64   `gorm:"primary_key;auto_increment" json:"id"`
	ParentID    uint64   `gorm:"default:0" json:"parent_id"`
	IsAnswer    bool     `gorm:"default:false" json:"is_answer"`
	UserID      *uint64  `gorm:"default:null" json:"user_id"`
	User        User     `gorm:"foreignKey:user_id" json:"user"`
	GuestName   string   `json:"guest_name"`
	GuestMobile string   `json:"guest_mobile"`
	Title       string   `json:"title"`
	Body        string   `json:"body"`
	GalleryID   *uint64  `gorm:"default:null" json:"gallery_id"`
	Gallery     *Gallery `gorm:"foreignKey:gallery_id" json:"gallery"`
	CreatedAt   string   `json:"created_at"`
}

func InitTicket(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Ticket{})
}

func (m *MysqlManager) CreateTicket(c *gin.Context, ctx context.Context, dto DTOs.CreateTicket, userID uint64) error {
	span, ctx := apm.StartSpan(ctx, "CreateTicket", "model")
	defer span.End()
	var parentTicket Ticket
	if dto.ParentID != 0 {
		err := m.GetConn().Model(&parentTicket).Where("id = ?", dto.ParentID).Update("is_answer", true).First(&parentTicket).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "خطا در دریافت تیکت ها",
				"error":   err.Error(),
				"type":    "model",
			})
			return err
		}
		if *parentTicket.UserID != userID && IsAdmin(c) == false {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "خطا در دریافت تیکت ها",
				"error":   "شما اجازه ارسال پاسخ به این تیکت را ندارید",
				"type":    "model",
			})
			return err
		}
	}

	ticket := Ticket{
		UserID: func() *uint64 {
			if userID == 0 {
				return nil
			}
			return &userID
		}(),
		ParentID:    dto.ParentID,
		IsAnswer:    dto.IsAnswer,
		GuestName:   dto.GuestName,
		GuestMobile: dto.GuestMobile,
		Title:       dto.Title,
		Body:        dto.Body,
		GalleryID: func() *uint64 {
			if dto.GalleryID == 0 {
				return nil
			}
			return &dto.GalleryID
		}(),
		CreatedAt: utils.NowTime(),
	}
	err := m.GetConn().Create(&ticket).Error
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

func (m *MysqlManager) GetAllTicketWithPagination(c *gin.Context, ctx context.Context, dto DTOs.IndexTicket, userID uint64) (*DTOs.Pagination, error) {
	span, ctx := apm.StartSpan(ctx, "GetAllTicketWithPagination", "model")
	defer span.End()
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

func (m *MysqlManager) GetTicketWithChildren(c *gin.Context, ctx context.Context, ticketID uint64) ([]Ticket, error) {
	span, ctx := apm.StartSpan(ctx, "GetTicketWithChildren", "model")
	defer span.End()
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
