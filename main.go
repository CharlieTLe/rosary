package main

import "fmt"
import "time"

func amen() {
	fmt.Printf("\nAmen.\n\n")
}
func makeSignOfTheCross() {
	fmt.Println("In the name of the Father, and of the Son, and of the Holy Spirit.")
	amen()
}

func sayApostolesCreed() {
	fmt.Println("I believe in God, the Father almighty, Creator of heaven and",
		"earth, and in Jesus Christ, his only Son, our Lord, who was conceived",
		"by the Holy Spirit, born of the Virgin Mary, suffered under Pontius Pilate,",
		"was crucified, died, and was buried; he descended into hell;",
		"on the third day he rose again from the dead; he ascended into heaven,",
		"and is seated at the right hand of God the Father almighty;",
		"from there He will come to judge the living and the dead. I believe in the Holy Spirit,",
		"the holy catholic Church, the communion of saints, the forgiveness of sins,",
		"the resurrection of the body, and life everlasting.")
	amen()
}

func sayOurFather() {
	fmt.Println("Our Father, who art in Heaven; hallowed by Thy name. Thy kingdom come,",
		"Thy will be done on earth as it is in Heaven. Give us this day our daily bread,",
		"and forgive us our trespasses as we forgive those who trespass against us,",
		"and lead us not into temptation; but deliver us from evil.")
	amen()
}

func hailMary(count int) {
	for i := 0; i < count; i++ {
		fmt.Println("Hail Mary, full of grace. The Lord is with thee. Blessed art thou among women,",
			"and blessed is the fruit of thy womb, Jesus. Holy Mary, Mother of God, pray for us sinners,",
			"now and at the hour of our death.")
		amen()
	}
}

func sayGloryBeToTheFather() {
	fmt.Println("Glory be to the Father, and to the Son, and to the Holy Spirit. " +
		"As it was in the beginning, is now, and ever shall be, world without end.")
	amen()
}

func announceMystery(ordinal int) {
	weekday := time.Now().Weekday()
	if weekday == time.Sunday || weekday == time.Wednesday {
		fmt.Println([]string{
			"Do not be amazed! You seek Jesus of Nazareth, the crucified. He has been raised; He is not here. Behold the place where they laid Him.",
			"So then the Lord Jesus, after He spoke to them, was taken up into heaven and took His seat at the right hand of God.",
			"And they were all filled with the holy Spirit and began to sepak in different tongues, as the Spirit enabled them to proclaim.",
			"You are the glory of Jerusalem! You are the great boast of our nation! You have done good things for Israel, and God is pleased with them. May the Almighty Lord bless you forever!",
			"A great sign appeared in the sky, a woman clothed with the sun, with the moon under her feet, and on her head a crown of twelve stars.",
		}[ordinal])
	} else if weekday == time.Monday || weekday == time.Saturday {
		fmt.Println([]string{
			`And when the angel had come to her, he said, "Hail, full of grace, the Lord is with you."`,
			`Elizabeth, filled with the holy Spirit, cried out in a loud voice and said, "MOst blessed are you among women, and blessed is the furit of your womb."`,
			"She gave birth to her firthborn Son. She wrapped Him in swaddling clothes and laid Him in a manger, because there was no room for them in the inn.",
			`When the days were completed for hteir purification according to the law of Moses, they took him up to Jerusalem to present Him to the Lord, just as it is written in the law of the Lord, " every male that opens the womb shall be consecrated to the Lord."`,
			"After three days they found Him in the temple, sitting in the midst of the teacer, listening to htem and asking them questions.",
		}[ordinal])
	} else if weekday == time.Tuesday || weekday == time.Friday {
		fmt.Println([]string{
			`He was in such agony and He prayed so fervently that His sweat became like grops of blood falling on the ground. When He rose from prayer and returned to His disciples, He found them sleeping from grief.`,
			`Then Pilate took Jesus and had Him scourged.`,
			"They stripped off His clothes and threw a scarlet military cloak about Him. Weaving a crown out of thorns, they placed it on His head, and a reed in His right hand.",
			`And carrying the cross Himself, He went out to what is called the Place of the Skull, in Hebrew, Golgotha.`,
			`Jesus cried out in a loud voice, "Father, into Your hands I commend My spirit"; and when He had said this He breathed His last.`,
		}[ordinal])
	} else if weekday == time.Thursday {
		fmt.Println([]string{
			`After Jesus was baptized, the heavens were opened [for Him]. and he saw the Spirit of God descending like a dove [and] coming upon Him. And a voice came from the heavens, saying, "This is My beloved Son, with whom I am well pleased."`,
			`His mother said to the servers, "Do whatever he tells you." Jesus told them, "Fill the jars with water." So they filled them to the brim.`,
			"As you go, make this proclamation: 'The kingdom of heaven is at hand.' Cure the sick, raise the dead, cleanse the lepers, drive out demons. Without cost you have received, without cost you are to give.",
			`While He was praying His face changed in appearance and His clothing became a dazzling white. Then from the cloud came a voice that said, "This is my chosen Son; listem to Him."`,
			`Then He took the bread, said the blessing, broke it, and gave it to them, saying, "This is My body, which will be given for you" And likewise the cup after they had eaten, saying, "This cup is the new covenenant in My blood."`,
		}[ordinal])
	}
	amen()
}

func hailHolyQueen() {
	fmt.Println("Hail, Holy Queen, Mother Mercy, our life, our sweetness and our hope, to thee do we cry, poor banished children of Eve;",
		"to thee do we send up our sighs, mourning and weeping in this valley of tears; turn, then, most gracious Advocate, thine eyes of",
		"mercy towards us, and after this, our exile, show unto us the blessed fruit of thy womb, Jesus. O clement, O loving, O sweet",
		"Virgin Mary! Pray for us O holy Mother of God, that we may be made worthy of the promises of Christ.")
}

func main() {
	makeSignOfTheCross()
	sayApostolesCreed()
	sayOurFather()
	hailMary(3)
	sayGloryBeToTheFather()
	for ordinal := 0; ordinal < 5; ordinal++ {
		announceMystery(ordinal)
		sayOurFather()
		hailMary(10)
		sayGloryBeToTheFather()
	}
	hailHolyQueen()
}
