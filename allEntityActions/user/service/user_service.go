package service

import (
	"trail1/allEntityActions/user"
	"trail1/entity"
)

type UserService struct {
	userRepo user.UserRepository
}

// NewUserService  returns a new UserService object
func NewUserService(userRepository user.UserRepository) user.UserService {
	return &UserService{userRepo: userRepository}
}

// Users returns all stored application users
func (us *UserService) Users() ([]entity.User, []error) {
	usrs, errs := us.userRepo.Users()
	if len(errs) > 0 {
		return nil, errs
	}
	return usrs, nil
}
func (us *UserService) Login(email string) (*entity.User, []error) {
	u, errs := us.userRepo.Login(email)
	if len(errs) > 0 {
		return nil, errs
	}
	return u, nil
}

// User retrieves an application user by its id
func (us *UserService) User(id uint) (*entity.User, []error) {
	usr, errs := us.userRepo.User(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, nil
}

// UpdateUser updates  a given application user
func (us *UserService) UpdateUser(user *entity.User) (*entity.User, []error) {
	usr, errs := us.userRepo.UpdateUser(user)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, nil
}

// DeleteUser deletes a given application user
func (us *UserService) DeleteUser(id uint) (*entity.User, []error) {
	usr, errs := us.userRepo.DeleteUser(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// StoreUser stores a given application user
func (us *UserService) StoreUser(user *entity.User) (*entity.User, []error) {
	usr, errs := us.userRepo.StoreUser(user)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, nil
}
func (us *UserService) ChangePassword(user *entity.User) (*entity.User, []error) {
	usr, errs := us.userRepo.ChangePassword(user)
	if len(errs) > 0 {
		return usr, errs
	}
	return usr, nil
}

// PhoneExists check if there is a user with a given phone number
func (us *UserService) PhoneExists(phone string) bool {
	exists := us.userRepo.PhoneExists(phone)
	return exists
}

// EmailExists checks if there exist a user with a given email address
func (us *UserService) EmailExists(email string) bool {
	exists := us.userRepo.EmailExists(email)
	return exists
}

// UserRoles returns list of roles a user has
func (us *UserService) UserRoles(user *entity.User) ([]entity.Role, []error) {
	userRoles, errs := us.userRepo.UserRoles(user)
	if len(errs) > 0 {
		return nil, errs
	}
	return userRoles, errs
}

func (ps *UserService) Properties() ([]entity.Property, []error) {
	pros, errs := ps.userRepo.Properties()
	if len(errs) > 0 {
		return nil, errs
	}
	return pros, errs
}

func (ps *UserService) Property(id uint) (*entity.Property, []error) {
	pro, errs := ps.userRepo.Property(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return pro, nil
}

func (ps *UserService) UserProperty(id uint) ([]entity.Property, []error) {
	pro, errs := ps.userRepo.UserProperty(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return pro, nil
}

func (ps *UserService) UserOrder(id uint) ([]entity.Order, []error) {
	ord, errs := ps.userRepo.UserOrder(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return ord, nil
}

func (ps UserService) UpdateProperty(property *entity.Property) (*entity.Property, []error) {
	pro, errs := ps.userRepo.UpdateProperty(property)
	if len(errs) > 0 {
		return nil, errs
	}
	return pro, nil
}

func (ps *UserService) DeleteProperty(id uint) (*entity.Property, []error) {
	pro, errs := ps.userRepo.DeleteProperty(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return pro, nil
}

func (ps *UserService) StoreProperty(property *entity.Property) (*entity.Property, []error) {
	pro, errs := ps.userRepo.StoreProperty(property)
	if len(errs) > 0 {
		return nil, errs
	}
	return pro, nil
}

func (ps *UserService) SearchProperty(index string) ([]entity.Property, error) {
	properties, err := ps.userRepo.SearchProperty(index)
	if err != nil {
		return nil, err
	}
	return properties, nil
}

func (ps *UserService) RateProperty(pro *entity.Property) (*entity.Property, []error) {
	prowithrate, err := ps.userRepo.RateProperty(pro)
	if err != nil {
		return prowithrate, err
	}
	return prowithrate, nil
}

func (ps *UserService) StorePropertyCateg(property *entity.Property) []error {
	err := ps.userRepo.StorePropertyCateg(property)
	if err != nil {
		return err
	}
	return nil
}

//func (ps *UserService) UserProperties(user *entity.User) ([]entity.Property, []error) {
//	errs, err := ps.userRepo.UserProperties(user)
//	if err != nil {
//		return nil,err
//	}
//	return errs, nil
//
//}




