class User{
    constructor(user){
        this.FirstName =user.FirstName
        this.LastName = user.LastName
        this.Email= user.Email
        this.Gender=user.Gender
        this.Role = user.Role
    }

    getRole(){
        return this.Role; 
    }

} export default User;