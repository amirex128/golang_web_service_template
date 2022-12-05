package models

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/brianvoe/gofakeit/v6"
	"go.elastic.co/apm/v2"
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
	if !manager.GetConn().Migrator().HasTable(&Ticket{}) {
		manager.GetConn().AutoMigrate(&Ticket{})
		for i := 0; i < 100; i++ {
			model := new(DTOs.CreateTicket)
			gofakeit.Struct(model)

			manager.CreateTicket(*model)
		}
	}

}

func (m *MysqlManager) CreateTicket(dto DTOs.CreateTicket) (*Ticket, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:CreateTicket", "model")
	defer span.End()
	var userID *uint64
	if dto.GuestMobile == "" {
		userID = GetUser(m.Ctx)
	}
	var parentTicket Ticket
	if dto.ParentID != 0 {
		err := m.GetConn().Model(&parentTicket).Where("id = ?", dto.ParentID).Update("is_answer", true).First(&parentTicket).Error
		if err != nil {
			return nil, errorx.New("خطا در دریافت تیکت ها", "model", err)
		}
		if *parentTicket.UserID != *userID && IsAdmin(m.Ctx) == false {
			return nil, errorx.New("خطا در دریافت تیکت ها", "model", err)
		}
	}

	ticket := &Ticket{
		UserID:      userID,
		ParentID:    dto.ParentID,
		IsAnswer:    false,
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
	err := m.GetConn().Create(ticket).Error
	if err != nil {
		return ticket, errorx.New("خطا در ثبت تیکت", "model", err)
	}
	return ticket, nil
}

func (m *MysqlManager) GetAllTicketWithPagination(dto DTOs.IndexTicket, userID uint64) (*DTOs.Pagination, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:GetAllTicketWithPagination", "model")
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
		return nil, errorx.New("خطا در دریافت تیکت ها", "model", err)
	}
	pagination.Data = tickets
	return pagination, nil
}

func (m *MysqlManager) GetTicketWithChildren(ticketID uint64) ([]Ticket, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:GetTicketWithChildren", "model")
	defer span.End()
	conn := m.GetConn()
	var tickets []Ticket
	var mainTicket Ticket
	err := conn.Where("id = ? ", ticketID).First(&mainTicket).Error
	if err != nil {
		return nil, errorx.New("خطا در دریافت تیکت ها", "model", err)
	}

	err = conn.Where("parent_id = ?", ticketID).Order("created_at").Preload("Gallery").Preload("User").Find(&tickets).Error
	if err != nil {
		return nil, errorx.New("خطا در دریافت تیکت ها", "model", err)
	}
	tickets = append([]Ticket{mainTicket}, tickets...)
	return tickets, nil
}

func (m *MysqlManager) DeleteTicket(ticketID uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showTicket", "model")
	defer span.End()
	ticket := Ticket{}
	err := m.GetConn().Where("id = ?", ticketID).First(&ticket).Error
	if err != nil {
		return errorx.New("تیکت یافت نشد", "model", err)
	}

	err = m.GetConn().Delete(&ticket).Error
	if err != nil {
		return errorx.New("خطا در حذف تیکت", "model", err)
	}
	return nil
}
