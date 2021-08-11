package user

import (
	"errors"
	"fmt"
	"github.com/freshpay/internal/config"
	"github.com/freshpay/internal/entities/OTP"
	"github.com/freshpay/internal/entities/user_management/user_session"
	"github.com/freshpay/internal/entities/user_management/wallet"
	utilities2 "github.com/freshpay/utilities"

	//"github.com/freshpay/internal/entities/user_management/wallet"
)

func VerifyPhoneNumber(phoneNumber string) bool {
	if phoneNumber == "1" {
		return false
	}
	return true
}

//SignUp will be used to create a user on signup
func SignUp(user *Detail) (err error) {
	phoneNumber := user.PhoneNumber
	if len(phoneNumber)!=10 || phoneNumber[0]=='0'{
		err=errors.New("phone number should be 10 digit long")
		return err
	}
	if !utilities2.IsNumeric(phoneNumber){
		err=errors.New("Phone number can contain characters 0-9")
		return err
	}
	/*
	   Make sure PhoneNumber doesn't exist
	*/
	var userTemp Detail
	err = GetUserByPhoneNumber(&userTemp, phoneNumber)
	if err == nil && userTemp.IsVerified {
		err = errors.New("Phone Number is already registered")
		return err
	} else if err == nil {
		err = DeleteUser(&userTemp)
		if err != nil {
			return err
		}
	}
	err = OTP.SendOTP(phoneNumber)
	if err != nil {
		return err
	}
	user.IsVerified=false
	user.ID = utilities2.CreateID(Prefix, IDLengthExcludingPrefix)

	/*
	 Encrypt the password
	 */
	var passwordHash string
	err= utilities2.GetEncryption(user.Password,&passwordHash)
	if err!=nil{
		return err
	}
	user.Password= passwordHash

	//now create the user
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}

	err = wallet.CreateWallet(user.ID)
	if err != nil {
		config.DB.Unscoped().Delete(&user)
		return err
	}
	return nil
}

//GetUserById will get the user infromation by using Id
func GetUserById(user *Detail, id string) (err error) {
	if err = config.DB.Where("ID = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByPhoneNumber will get the user details by phone number
func GetUserByPhoneNumber(user *Detail, phoneNumber string) (err error) {
	if err = config.DB.Where("phone_number = ?", phoneNumber).First(user).Error; err != nil {
		return err
	}
	return nil
}

//update the user
func UpdateUser(user *Detail) (err error) {
	err = config.DB.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

//Delete a user
func DeleteUser(user *Detail) (err error) {
	err = config.DB.Where("id = ?", user.ID).Delete(user).Error
	return err
}

//Login will login the user and will create a user_session
func LoginByPassword(phoneNumber string, password string, Session *user_session.Detail, user *Detail) (err error) {
	err = GetUserByPhoneNumber(user, phoneNumber)
	if err == nil {
		if !user.IsVerified {
			err = errors.New("Phone Number is not verified, please signup again")
			fmt.Println(err)
			/*
			   need to remove this line
			*/
			return err
		}
		if  !utilities2.MatchPassword(password,user.Password){
			err = errors.New("Password is Wrong")
		} else {
			err=user_session.GetActiveSessionByUserId(Session,user.ID)
			if err==nil{
				return nil
			}
			Session.UserId = user.ID
			err = user_session.CreateSession(Session)
		}
	}else{
		err=errors.New("Phone Number is wrong or not registered")
	}
	return err
}

//set verified user by phone number
func SetVerifiedUserByPhoneNumber(phoneNumber string) (err error) {
	var user Detail
	err = GetUserByPhoneNumber(&user, phoneNumber)
	if err == nil {
		user.IsVerified = true
		err = UpdateUser(&user)
	}
	return err
}
