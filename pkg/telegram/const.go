package telegram

const (
	StepStart = iota + 1
	StepManNik
	StepManConfirmation
	StepManWait

	StepWomanData
	StepWomanCorrection
	StepWomanNik
	StepWomanConfirmation

	startMessage = "Добро пожаловать!\n" +
		"Введите пожалйста свой пол"
	helpMessage   = "Чтоб перезапустить бота пропишите /start или нажмете на команду"
	unknownGender = "Попробуйте ввести заного, я не понял вашего пола"
	unknownNik    = "Такой ник не может существовать!\n" +
		"Пожалуйста попробуйте снова"

	stepManNik = "Отлично\nТеперь введи ник твоей девушки"

	stepWomanData       = "Введи дату начала последних месячных в форме 01.01.0001"
	stepWomanCorrection = "Отлично! Я запомнил!\nВведи сколько дней они у тебя дляться"
	stepWomanNik        = "Отлично\nТеперь введи ник твоего парня"

	stepManConfirmation   = "Ожидайте пока ваша девушка подтвуредить, что вы её парень"
	stepWomanConfirmation = "Ожидайте пока ваш парень подтвуредить, что вы его девушка"

	startCommand   = "start"
	helpCommand    = "help"
	unknownCommand = "Я не знаю такой команды!"
)
