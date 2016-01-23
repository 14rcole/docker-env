package errors

func InputError(message string){
	fmt.Println("Oops! Looks like there's somthing wrong with your input");
	fmt.Println("The error occured at %s", message);
}
