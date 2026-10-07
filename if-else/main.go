package main

import "fmt"

func main() {

	// age ভেরিয়েবল ইনিশিয়ালাইজ করা হচ্ছে
	age := 70
	
	// সাধারণ if-else কন্ডিশন ব্যবহার করে বয়স চেক করা হচ্ছে
	if age < 18 {
		fmt.Println("You are a minor.") // বয়স ১৮ এর নিচে হলে
	} else if age >= 18 && age < 65 {
		fmt.Println("You are an adult.") // বয়স ১৮ থেকে ৬৪ এর মধ্যে হলে
	} else {
		fmt.Println("You are a senior citizen.") // অন্যথায় (বয়স ৬৫ বা তার বেশি হলে)
	}

	// if-else scope এর উদাহরণ
	score := 85
	// এই score ভেরিয়েবলটি বাইরের স্কোপের, তাই এটি যেকোনো জায়গা থেকে ব্যবহার করা যাবে
	fmt.Println("outside if-else scope, score is:", score)

	// if কন্ডিশনের ভেতরে নতুন একটি score ভেরিয়েবল (score:=50) ডিক্লেয়ার করা হয়েছে।
	// এই score ভেরিয়েবলটির অস্তিত্ব শুধুমাত্র এই if-else ব্লকের ভেতরেই সীমাবদ্ধ থাকবে।
	if score:=50; score >= 90 { 
		fmt.Println("You got an A!")
	} else if score >= 80 {
		fmt.Println("You got a B!")
	} else if score >= 70 {
		fmt.Println("You got a C!")
	} else if score >= 60 {
		fmt.Println("You got a D!")
	} else {
		fmt.Println("You failed.")
	}

}