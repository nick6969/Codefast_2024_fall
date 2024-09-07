package controller

import (
	"codefast_2024/app"
	"io/fs"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
)

type answer struct {
	QuestionID string `json:"question_id"`
	Answer     struct {
		OptionID string `json:"option_id"`
		Value    string `json:"value"`
	} `json:"answer"`
}

func AnswerLabor(ctx *gin.Context) {
	var answers []answer

	if err := ctx.BindJSON(&answers); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	checkAnswers := laborAnswers()

	result := ""

	for _, a := range answers {
		check := sliceFirst(checkAnswers, func(ca laborAnswer) bool {
			return ca.QuestionID == a.QuestionID
		})

		if check == nil {
			ctx.JSON(400, gin.H{"error": "Invalid question ID"})
			return
		}

		if slices.Contains(check.WrongAnswerIDs, a.Answer.OptionID) {
			result += check.Description + "<br>"
		}
	}

	ctx.JSON(200, gin.H{"message": result})
}

func laborAnswers() []laborAnswer {
	return []laborAnswer{
		{
			QuestionID:     "1",
			WrongAnswerIDs: []string{"2"},
			Description:    "勞動基準法施行細則第7條：「勞動契約應依本法有關規定約定下列事項：一、工作場所及應從事之工作。二、工作開始與終止之時間、休息時間、休假、例假、休息日、請假及輪班制之換班。三、工資之議定、調整、計算、結算與給付之日期及方法。四、勞動契約之訂定、終止及退休。五、資遣費、退休金、其他津貼及獎金。六、勞工應負擔之膳宿費及工作用具費。七、安全衛生。八、勞工教育及訓練。九、福利。十、災害補償及一般傷病補助。十一、應遵守之紀律。十二、獎懲。十三、其他勞資權利義務有關事項。」",
		},
		{
			QuestionID:     "2",
			WrongAnswerIDs: []string{"1"},
			Description:    "勞動基準法第21條規定，工資由勞雇雙方議定之。但不得低於基本工資。故勞工每月在正常工作時間內所得之工資，不得低於每月基本工資扣除因請假而未發之每日基本工資後之餘額。",
		},
		{
			QuestionID:     "3",
			WrongAnswerIDs: []string{"3"},
			Description:    "勞動基準法施行細則第14條之1:「雇主提供之前項明細，得以紙本、電子資料傳輸方式或其他勞工可隨時取得及得列印之資料為之。」，沒有固定時間也是可以的",
		},
		{
			QuestionID:     "4",
			WrongAnswerIDs: []string{"2", "3"},
			Description:    "勞動基準法第23條第1項規定，「工資給付，除非當事人有特別約定或是按月預付，每月至少定期發給2次，並且應提供工資個項目計算方式明細；按件計酬者亦同。」",
		},
		{
			QuestionID:     "5",
			WrongAnswerIDs: []string{"1", "3", "4"},
			Description:    "勞動基準法第43條：「勞工因婚、喪、疾病或其他正當事由得請假；請假應給之假期及事假以外期間內工資給付之最低標準，由中央主管機關定之。」「請假為勞工權利」，且雇主不能拒絕勞工依法行使權利的結果",
		},
		{
			QuestionID:     "6",
			WrongAnswerIDs: []string{"2", "3"},
			Description:    "勞基法第22條第2項規定，必須將薪資全額給付給勞工，員工處罰，處罰的理由要有足夠正當性，亦不得由薪水直接扣除。",
		},
		{
			QuestionID:     "7",
			WrongAnswerIDs: []string{},
			Description:    "",
		},
		{
			QuestionID:     "8",
			WrongAnswerIDs: []string{},
			Description:    "",
		},
		{
			QuestionID:     "9",
			WrongAnswerIDs: []string{},
			Description:    "",
		},
		{
			QuestionID:     "10",
			WrongAnswerIDs: []string{},
			Description:    "",
		},
	}
}

type laborAnswer struct {
	QuestionID     string   `json:"question_id"`
	WrongAnswerIDs []string `json:"wrong_answer_ids"`
	Description    string   `json:"description"`
}

func sliceFirst[T any](slice []T, f func(T) bool) *T {
	for _, element := range slice {
		if f(element) {
			return &element
		}
	}

	return nil
}

func Lunar(app *app.App) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		now := time.Now().In(time.FixedZone("Asia/Taipei", 8*60*60))

		dateString := now.Format("20060102")

		file, err := getLunarImage(dateString, app)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "Internal server error"})
			return
		}

		ctx.Data(200, "image/png", file)
	}

}

func getLunarImage(date string, app *app.App) ([]byte, error) {
	file, err := fs.ReadFile(app.PageFS, date+".png")
	if err != nil {
		return nil, err
	}

	return file, nil
}
