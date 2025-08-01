// Code generated by thriftgo (0.4.1). DO NOT EDIT.

package user

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type UserInfoDetail struct {
	UserID    *string `thrift:"user_id,1,optional" frugal:"1,optional,string" form:"user_id" json:"user_id,omitempty" query:"user_id"`
	Name      *string `thrift:"name,11,optional" frugal:"11,optional,string" form:"name" json:"name,omitempty" query:"name"`
	NickName  *string `thrift:"nick_name,12,optional" frugal:"12,optional,string" form:"nick_name" json:"nick_name,omitempty" query:"nick_name"`
	AvatarURL *string `thrift:"avatar_url,13,optional" frugal:"13,optional,string" form:"avatar_url" json:"avatar_url,omitempty" query:"avatar_url"`
	Email     *string `thrift:"email,14,optional" frugal:"14,optional,string" form:"email" json:"email,omitempty" query:"email"`
	Mobile    *string `thrift:"mobile,15,optional" frugal:"15,optional,string" form:"mobile" json:"mobile,omitempty" query:"mobile"`
}

func NewUserInfoDetail() *UserInfoDetail {
	return &UserInfoDetail{}
}

func (p *UserInfoDetail) InitDefault() {
}

var UserInfoDetail_UserID_DEFAULT string

func (p *UserInfoDetail) GetUserID() (v string) {
	if p == nil {
		return
	}
	if !p.IsSetUserID() {
		return UserInfoDetail_UserID_DEFAULT
	}
	return *p.UserID
}

var UserInfoDetail_Name_DEFAULT string

func (p *UserInfoDetail) GetName() (v string) {
	if p == nil {
		return
	}
	if !p.IsSetName() {
		return UserInfoDetail_Name_DEFAULT
	}
	return *p.Name
}

var UserInfoDetail_NickName_DEFAULT string

func (p *UserInfoDetail) GetNickName() (v string) {
	if p == nil {
		return
	}
	if !p.IsSetNickName() {
		return UserInfoDetail_NickName_DEFAULT
	}
	return *p.NickName
}

var UserInfoDetail_AvatarURL_DEFAULT string

func (p *UserInfoDetail) GetAvatarURL() (v string) {
	if p == nil {
		return
	}
	if !p.IsSetAvatarURL() {
		return UserInfoDetail_AvatarURL_DEFAULT
	}
	return *p.AvatarURL
}

var UserInfoDetail_Email_DEFAULT string

func (p *UserInfoDetail) GetEmail() (v string) {
	if p == nil {
		return
	}
	if !p.IsSetEmail() {
		return UserInfoDetail_Email_DEFAULT
	}
	return *p.Email
}

var UserInfoDetail_Mobile_DEFAULT string

func (p *UserInfoDetail) GetMobile() (v string) {
	if p == nil {
		return
	}
	if !p.IsSetMobile() {
		return UserInfoDetail_Mobile_DEFAULT
	}
	return *p.Mobile
}
func (p *UserInfoDetail) SetUserID(val *string) {
	p.UserID = val
}
func (p *UserInfoDetail) SetName(val *string) {
	p.Name = val
}
func (p *UserInfoDetail) SetNickName(val *string) {
	p.NickName = val
}
func (p *UserInfoDetail) SetAvatarURL(val *string) {
	p.AvatarURL = val
}
func (p *UserInfoDetail) SetEmail(val *string) {
	p.Email = val
}
func (p *UserInfoDetail) SetMobile(val *string) {
	p.Mobile = val
}

var fieldIDToName_UserInfoDetail = map[int16]string{
	1:  "user_id",
	11: "name",
	12: "nick_name",
	13: "avatar_url",
	14: "email",
	15: "mobile",
}

func (p *UserInfoDetail) IsSetUserID() bool {
	return p.UserID != nil
}

func (p *UserInfoDetail) IsSetName() bool {
	return p.Name != nil
}

func (p *UserInfoDetail) IsSetNickName() bool {
	return p.NickName != nil
}

func (p *UserInfoDetail) IsSetAvatarURL() bool {
	return p.AvatarURL != nil
}

func (p *UserInfoDetail) IsSetEmail() bool {
	return p.Email != nil
}

func (p *UserInfoDetail) IsSetMobile() bool {
	return p.Mobile != nil
}

func (p *UserInfoDetail) Read(iprot thrift.TProtocol) (err error) {
	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 11:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField11(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 12:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField12(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 13:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField13(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 14:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField14(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 15:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField15(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UserInfoDetail[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UserInfoDetail) ReadField1(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.UserID = _field
	return nil
}
func (p *UserInfoDetail) ReadField11(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.Name = _field
	return nil
}
func (p *UserInfoDetail) ReadField12(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.NickName = _field
	return nil
}
func (p *UserInfoDetail) ReadField13(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.AvatarURL = _field
	return nil
}
func (p *UserInfoDetail) ReadField14(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.Email = _field
	return nil
}
func (p *UserInfoDetail) ReadField15(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.Mobile = _field
	return nil
}

func (p *UserInfoDetail) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("UserInfoDetail"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField11(oprot); err != nil {
			fieldId = 11
			goto WriteFieldError
		}
		if err = p.writeField12(oprot); err != nil {
			fieldId = 12
			goto WriteFieldError
		}
		if err = p.writeField13(oprot); err != nil {
			fieldId = 13
			goto WriteFieldError
		}
		if err = p.writeField14(oprot); err != nil {
			fieldId = 14
			goto WriteFieldError
		}
		if err = p.writeField15(oprot); err != nil {
			fieldId = 15
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UserInfoDetail) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetUserID() {
		if err = oprot.WriteFieldBegin("user_id", thrift.STRING, 1); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.UserID); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}
func (p *UserInfoDetail) writeField11(oprot thrift.TProtocol) (err error) {
	if p.IsSetName() {
		if err = oprot.WriteFieldBegin("name", thrift.STRING, 11); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.Name); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 11 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 11 end error: ", p), err)
}
func (p *UserInfoDetail) writeField12(oprot thrift.TProtocol) (err error) {
	if p.IsSetNickName() {
		if err = oprot.WriteFieldBegin("nick_name", thrift.STRING, 12); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.NickName); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 12 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 12 end error: ", p), err)
}
func (p *UserInfoDetail) writeField13(oprot thrift.TProtocol) (err error) {
	if p.IsSetAvatarURL() {
		if err = oprot.WriteFieldBegin("avatar_url", thrift.STRING, 13); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.AvatarURL); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 13 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 13 end error: ", p), err)
}
func (p *UserInfoDetail) writeField14(oprot thrift.TProtocol) (err error) {
	if p.IsSetEmail() {
		if err = oprot.WriteFieldBegin("email", thrift.STRING, 14); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.Email); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 14 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 14 end error: ", p), err)
}
func (p *UserInfoDetail) writeField15(oprot thrift.TProtocol) (err error) {
	if p.IsSetMobile() {
		if err = oprot.WriteFieldBegin("mobile", thrift.STRING, 15); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.Mobile); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 15 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 15 end error: ", p), err)
}

func (p *UserInfoDetail) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserInfoDetail(%+v)", *p)

}

func (p *UserInfoDetail) DeepEqual(ano *UserInfoDetail) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.UserID) {
		return false
	}
	if !p.Field11DeepEqual(ano.Name) {
		return false
	}
	if !p.Field12DeepEqual(ano.NickName) {
		return false
	}
	if !p.Field13DeepEqual(ano.AvatarURL) {
		return false
	}
	if !p.Field14DeepEqual(ano.Email) {
		return false
	}
	if !p.Field15DeepEqual(ano.Mobile) {
		return false
	}
	return true
}

func (p *UserInfoDetail) Field1DeepEqual(src *string) bool {

	if p.UserID == src {
		return true
	} else if p.UserID == nil || src == nil {
		return false
	}
	if strings.Compare(*p.UserID, *src) != 0 {
		return false
	}
	return true
}
func (p *UserInfoDetail) Field11DeepEqual(src *string) bool {

	if p.Name == src {
		return true
	} else if p.Name == nil || src == nil {
		return false
	}
	if strings.Compare(*p.Name, *src) != 0 {
		return false
	}
	return true
}
func (p *UserInfoDetail) Field12DeepEqual(src *string) bool {

	if p.NickName == src {
		return true
	} else if p.NickName == nil || src == nil {
		return false
	}
	if strings.Compare(*p.NickName, *src) != 0 {
		return false
	}
	return true
}
func (p *UserInfoDetail) Field13DeepEqual(src *string) bool {

	if p.AvatarURL == src {
		return true
	} else if p.AvatarURL == nil || src == nil {
		return false
	}
	if strings.Compare(*p.AvatarURL, *src) != 0 {
		return false
	}
	return true
}
func (p *UserInfoDetail) Field14DeepEqual(src *string) bool {

	if p.Email == src {
		return true
	} else if p.Email == nil || src == nil {
		return false
	}
	if strings.Compare(*p.Email, *src) != 0 {
		return false
	}
	return true
}
func (p *UserInfoDetail) Field15DeepEqual(src *string) bool {

	if p.Mobile == src {
		return true
	} else if p.Mobile == nil || src == nil {
		return false
	}
	if strings.Compare(*p.Mobile, *src) != 0 {
		return false
	}
	return true
}
