package model

type User struct {
	ID          string `json:"id"`
	UID         string `json:"uid"`
	Nickname    string `bson:"nick_name" json:"nick_name"`
	FaceURL     string `bson:"face_url" json:"face_url"`
	Gender      int32  `bson:"gender" json:"gender"`
	Email       string `bson:"email" json:"email"`
	PhoneNumber string `bson:"phone_number" json:"phone_number"`
	Birth       int64  `bson:"birth" json:"birth"`
	Status      int32  `bson:"status" json:"status"`
	//Password    string    `bson:"password" json:"password"`
	Version int `bson:"version" json:"version"`
}
