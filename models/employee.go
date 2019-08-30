package models

import (
	"ncd_operators/pkg/raven"
	s "ncd_operators/pkg/settings"
	"os"
	"time"
	// "fmt"
)

type Employee struct {
	Created           *time.Time `json:"created"`
	Updated           *time.Time `json:"updated"`
	Id                int        `json:"id"`
	SendSMS           bool       `json:"send_sms" db:"send_sms"`
	SendEmail         bool       `json:"send_email" db:"send_email"`
	FullNameRu        string     `json:"full_name_ru" db:"full_name_ru"`
	FullNameEn        string     `json:"full_name_en" db:"full_name_en"`
	PassportSerial    string     `json:"passport_serial" db:"passport_serial"`
	PassportGivenDate string     `json:"passport_given_date" db:"passport_given_date"`
	BirthDate         string     `json:"birth_date" db:"birth_date"`
	BirthPlaceRu      string     `json:"birth_place_ru" db:"birth_place_ru"`
	BirthPlaceEn      string     `json:"birth_place_en" db:"birth_place_en"`
	LivingAddressRu   string     `json:"living_address_ru" db:"living_address_ru"`
	LivingAddressEn   string     `json:"living_address_en" db:"living_address_en"`
	Gender            string     `json:"gender" db:"gender"`
	Inn               string     `json:"inn" db:"inn"`
	Phone             string     `json:"phone" db:"phone"`
	Email             string     `json:"email" db:"email"`
	RegisterNumber    string     `json:"register_number" db:"register_number"`
}

func (e Employee) CreatedFormat() string {
	return e.Created.Format("02.01.2006 15:04:05")
}
func EmployeeById(id string) (*Employee, error) {
	var (
		err error
		e   Employee
	)
	err = s.UDB.Get(&e, `SELECT e.id, e.register_number, e.passport_serial FROM employee e WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

// EDUCATION
type Education struct {
	EducationType    `db:"education_type"`
	Id               int             `json:"id"`
	NameRu           string          `json:"name_ru" db:"name_ru"`
	NameEn           string          `json:"name_en" db:"name_en"`
	AddressRu        string          `json:"address_ru" db:"address_ru"`
	AddressEn        string          `json:"address_en" db:"address_en"`
	SpecializationRu string          `json:"specialization_ru" db:"specialization_ru"`
	SpecializationEn string          `json:"specialization_en" db:"specialization_en"`
	DateStarted      *time.Time      `json:"date_started" db:"date_started"`
	DateFinished     *time.Time      `json:"date_finished" db:"date_finished"`
	AdditionalRu     string          `json:"additional_ru" db:"additional_ru"`
	AdditionalEn     string          `json:"additional_en" db:"additional_en"`
	IsNew            bool            `json:"is_new" db:"is_new"`
	Files            []EducationFile `db:"files"`
}

func (e Education) DateStartedFormat() string {
	return e.DateStarted.Format("02.01.2006")
}

func (e Education) DateFinishedFormat() string {
	return e.DateStarted.Format("02.01.2006")
}

type EducationFile struct {
	File string `db:"file"`
}

func (ef *EducationFile) Url() string {
	return "/media/" + ef.File
}

func (ef *EducationFile) Size() int64 {
	f, err := os.Stat(s.Conf["MEDIA_PATH"] + "/" + ef.File)
	if err != nil {
		return 0
	}
	var bytes int64
	var kilobytes int64
	bytes = f.Size()
	kilobytes = bytes / 1024
	return kilobytes
}

// ARMY
type Army struct {
	Id               int        `json:"id"`
	NameRu           string     `json:"name_ru" db:"name_ru"`
	NameEn           string     `json:"name_en" db:"name_en"`
	RegionRu         string     `json:"region_ru" db:"region_ru"`
	RegionEn         string     `json:"region_en" db:"region_en"`
	SpecializationRu string     `json:"specialization_ru" db:"specialization_ru"`
	SpecializationEn string     `json:"specialization_en" db:"specialization_en"`
	DateStarted      *time.Time `json:"date_started" db:"date_started"`
	DateFinished     *time.Time `json:"date_finished" db:"date_finished"`
	RankRu           string     `json:"rank_ru" db:"rank_ru"`
	RankEn           string     `json:"rank_en" db:"rank_en"`
	IsNew            bool       `json:"is_new" db:"is_new"`
	Files            []ArmyFile
}

func (a Army) DateStartedFormat() string {
	if a.DateStarted != nil {
		return a.DateStarted.Format("02.01.2006")
	}
	return ""
}

func (a Army) DateFinishedFormat() string {
	if a.DateFinished != nil {
		return a.DateFinished.Format("02.01.2006")
	}
	return ""
}

type ArmyFile struct {
	File string `db:"file"`
}

func (ef *ArmyFile) Url() string {
	return "/media/" + ef.File
}

func (ef *ArmyFile) Size() int64 {
	f, err := os.Stat(s.Conf["MEDIA_PATH"] + "/" + ef.File)
	if err != nil {
		raven.ReportIfError(err)
		return 0
	}
	var bytes int64
	var kilobytes int64
	bytes = f.Size()
	kilobytes = bytes / 1024
	return kilobytes
}

// LANGUAGE
type Language struct {
	DLanguage         `db:"dlanguage"`
	Id                int  `json:"id"`
	Level             int  `json:"level" db:"level"`
	IsRequiredToCheck bool `json:"is_required_to_check" db:"is_required_to_check"`
	IsNew             bool `json:"is_new" db:"is_new"`
	Files             []LanguageFile
}

func (l *Language) LevelDisplay() string {
	if l.Level == 1 {
		return "Элементарный"
	} else if l.Level == 2 {
		return "Средний"
	} else if l.Level == 3 {
		return "Продвинутый"
	}
	return "Не выбран уровен"
}

type LanguageFile struct {
	File string `db:"file"`
}

func (ef *LanguageFile) Url() string {
	return "/media/" + ef.File
}

func (ef *LanguageFile) Size() int64 {
	f, err := os.Stat(s.Conf["MEDIA_PATH"] + "/" + ef.File)
	if err != nil {
		raven.ReportIfError(err)
		return 0
	}
	var bytes int64
	var kilobytes int64
	bytes = f.Size()
	kilobytes = bytes / 1024
	return kilobytes
}

// FAMILY
type Family struct {
	Id             int  `json:"id"`
	Status         int  `json:"status" db:"status"`
	ChildrenAmount int  `json:"children_amount" db:"children_amount"`
	IsNew          bool `json:"is_new" db:"is_new"`
	Files          []FamilyFile
}

func (f *Family) StatusDisplay() string {
	switch f.Status {
	case 0:
		return "Не указан"
	case 1:
		return "Не женат/не замужем"
	case 2:
		return "Женат/Замужем"
	case 3:
		return "Разведен/Разведена"
	case 4:
		return "Вдовец/Вдова"
	case 5:
		return "Второй брак"
	case 6:
		return "Гражданский брак"
	default:
		return "Не указан"
	}
}

type FamilyFile struct {
	File string `db:"file"`
}

func (ef *FamilyFile) Url() string {
	return "/media/" + ef.File
}

func (ef *FamilyFile) Size() int64 {
	f, err := os.Stat(s.Conf["MEDIA_PATH"] + "/" + ef.File)
	if err != nil {
		raven.ReportIfError(err)
		return 0
	}
	var bytes int64
	var kilobytes int64
	bytes = f.Size()
	kilobytes = bytes / 1024
	return kilobytes
}

// REWARD
type Reward struct {
	Id            int    `json:"id"`
	NameRu        string `json:"name_ru" db:"name_ru"`
	NameEn        string `json:"name_en" db:"name_en"`
	DescriptionRu string `json:"description_ru" db:"description_ru"`
	DescriptionEn string `json:"description_en" db:"description_en"`
	IsNew         bool   `json:"is_new" db:"is_new"`
	Files         []RewardFile
}

type RewardFile struct {
	File string `db:"file"`
}

func (ef *RewardFile) Url() string {
	return "/media/" + ef.File
}

func (ef *RewardFile) Size() int64 {
	f, err := os.Stat(s.Conf["MEDIA_PATH"] + "/" + ef.File)
	if err != nil {
		raven.ReportIfError(err)
		return 0
	}
	var bytes int64
	var kilobytes int64
	bytes = f.Size()
	kilobytes = bytes / 1024
	return kilobytes
}

// RELATIVE
type Relative struct {
	Id               int        `json:"id"`
	Level            int        `json:"level" db:"level"`
	FullNameRu       string     `json:"full_name_ru" db:"fullname_ru"`
	FullNameEn       string     `json:"full_name_en" db:"fullname_en"`
	BirthDate        *time.Time `json:"birth_date" db:"birth_date"`
	StudyWorkPlaceRu string     `json:"study_work_place_ru" db:"study_work_place_ru"`
	StudyWorkPlaceEn string     `json:"study_work_place_en" db:"study_work_place_en"`
	PositionRu       string     `json:"position_ru" db:"position_ru"`
	PositionEn       string     `json:"position_en" db:"position_en"`
	LivingAddressRu  string     `json:"living_address_ru" db:"living_address_ru"`
	LivingAddressEn  string     `json:"living_address_en" db:"living_address_en"`
	IsNew            bool       `json:"is_new" db:"is_new"`
	Files            []RelativeFile
}

func (r Relative) BirthDateFormat() string {
	if r.BirthDate != nil {
		return r.BirthDate.Format("02.01.2006")
	}
	return ""
}
func (r Relative) LevelDisplay() string {
	switch r.Level {
	case 1:
		return "Жена"
	case 2:
		return "Муж"
	case 3:
		return "Мать"
	case 4:
		return "Отец"
	case 5:
		return "Сын"
	case 6:
		return "Дочь"
	case 7:
		return "Бабушка"
	case 8:
		return "Дедушка"
	default:
		return ""
	}
}

type RelativeFile struct {
	File string `db:"file"`
}

func (ef *RelativeFile) Url() string {
	return "/media/" + ef.File
}

func (ef *RelativeFile) Size() int64 {
	f, err := os.Stat(s.Conf["MEDIA_PATH"] + "/" + ef.File)
	if err != nil {
		raven.ReportIfError(err)
		return 0
	}
	var bytes int64
	var kilobytes int64
	bytes = f.Size()
	kilobytes = bytes / 1024
	return kilobytes
}

// EXPERIENCE
type Experience struct {
	Id             int        `json:"id"`
	OrganizationRu string     `json:"organization_ru" db:"organization_ru"`
	OrganizationEn string     `json:"organization_en" db:"organization_en"`
	DateStarted    *time.Time `json:"date_started" db:"date_started"`
	DateFinished   *time.Time `json:"date_finished" db:"date_finished"`
	PositionRu     string     `json:"position_ru" db:"position_ru"`
	PositionEn     string     `json:"position_en" db:"position_en"`
	SubDivisionRu  string     `json:"sub_division_ru" db:"sub_division_ru"`
	SubDivisionEn  string     `json:"sub_division_en" db:"sub_division_en"`
	AddressRu      string     `json:"address_ru" db:"address_ru"`
	AddressEn      string     `json:"address_en" db:"address_en"`
	IsNew          bool       `json:"is_new" db:"is_new"`
	Files          []ExperienceFile
}

func (e Experience) DateStartedFormat() string {
	if e.DateStarted != nil {
		return e.DateStarted.Format("02.01.2006")
	}
	return ""
}

func (e Experience) DateFinishedFormat() string {
	if e.DateFinished != nil {
		return e.DateFinished.Format("02.01.2006")
	}
	return ""
}

type ExperienceFile struct {
	File string `db:"file"`
}

func (ef *ExperienceFile) Url() string {
	return "/media/" + ef.File
}

func (ef *ExperienceFile) Size() int64 {
	f, err := os.Stat(s.Conf["MEDIA_PATH"] + "/" + ef.File)
	if err != nil {
		raven.ReportIfError(err)
		return 0
	}
	var bytes int64
	var kilobytes int64
	bytes = f.Size()
	kilobytes = bytes / 1024
	return kilobytes
}
