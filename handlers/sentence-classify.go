package handlers

import (
	"regexp"
	"strings"
)

// SentenceClassify 句子分类
func SentenceClassify(sentence string) bool {
	// 检查输入是否为空值
	if len(strings.TrimSpace(sentence)) == 0 {
		return false
	}

	// 常见的中文疑问词集合
	//chineseQuestionWords := map[string]bool{
	//	"吗": true, "么": true, "什么": true, "怎么": true, "为什么": true, "是否": true, "哪": true,
	//	"几": true, "多少": true, "多大": true, "谁": true, "啥": true, "哪儿": true, "能否": true,
	//	"哪里": true, "哪个": true, "何时": true, "怎样": true, "咋样": true, "有何": true, "有么": true,
	//	"对吧": true, "好吗": true, "如何": true, "为啥": true, "难道": true, "有没有": true,
	//}

	// 常见英文疑问词集合
	englishQuestionWords := map[string]bool{
		"what": true, "how": true, "why": true, "is": true, "are": true, "does": true,
		"do": true, "did": true, "can": true, "could": true, "will": true, "would": true,
		"shall": true, "should": true, "who": true, "where": true, "when": true,
		"which": true, "whom": true,
	}

	// 常见的中英文问句短语
	questionPhrases := []string{
		"你觉得呢", "应该可以吧", "你认为呢", "行不行", "是不是", "可以吗", "能不能", "好不好",
		"会不会", "这样行吗", "可以不", "对不对", "难道不", "该如何", "怎么办", "这样不好吧",
		"可不可以", "你觉得", "你怎么看", "行了吧", "aren't you", "isn't it", "could it be", "how about",
	}

	// 过滤侮辱性或无意义的短语
	nonQuestionPhrases := []string{
		"你妈的", "去死", "傻逼", "你丫的", "他妈的", "草你妈", "日你妈", "傻b", "智障",
		"混蛋", "混账", "滚蛋", "fuck", "shit", "idiot",
	}

	// 转换为小写字母进行检查（英文部分）
	lowerSentence := strings.ToLower(sentence)

	// 检查是否包含侮辱性或无意义的短语
	for _, phrase := range nonQuestionPhrases {
		if strings.Contains(lowerSentence, phrase) {
			return false
		}
	}

	// 检查是否包含常见的问句短语
	for _, phrase := range questionPhrases {
		if strings.Contains(lowerSentence, phrase) {
			return true
		}
	}

	// 检查是否以数字或纯标点符号作为主要内容的问句
	if matched, _ := regexp.MatchString(`^\d+[?？]$`, sentence); matched {
		return false
	}
	if matched, _ := regexp.MatchString(`^[?？]+$`, sentence); matched {
		return false
	}

	// 检查中文问句
	chinesePattern := regexp.MustCompile(`(吗|么|什么|怎么|为什么|是否|哪|几|多少|多大|谁|啥|哪儿|能否|哪里|哪个|何时|怎样|咋样|有何|有么|对吧|好吗|如何|为啥|难道|有没有)|(?:.*[啊呢吗呀]?[?？])`)
	if chinesePattern.MatchString(sentence) {
		return true
	}

	// 针对含有方向、比较等关键词的句子进行特殊处理
	directionWords := []string{"方向", "趋势", "前景", "可能", "选择", "如何", "更好", "更优"}
	for _, word := range directionWords {
		if strings.Contains(sentence, word) && (strings.Contains(sentence, "那些") || strings.Contains(sentence, "哪个") || strings.Contains(sentence, "哪种")) {
			return true
		}
	}

	// 检查隐含疑问语气
	hiddenQuestionPatterns := []string{`.*了没有$`, `.*了没$`, `.*吗$`, `有没有.*`, `.*咋.*`, `.*行不.*`}
	for _, pattern := range hiddenQuestionPatterns {
		if matched, _ := regexp.MatchString(pattern, sentence); matched {
			return true
		}
	}

	// 专门处理'是吧'、'不是'等情况
	if strings.HasSuffix(sentence, "是吧") || strings.HasSuffix(sentence, "吧") || strings.HasSuffix(sentence, "不是") {
		return true
	}

	// 更准确的问号结尾检查
	trimmedSentence := strings.TrimSpace(sentence)
	if strings.HasSuffix(trimmedSentence, "？") || strings.HasSuffix(trimmedSentence, "?") {
		return true
	}

	// 额外处理问号前后的问句结构
	precedingText := strings.TrimSpace(strings.TrimSuffix(trimmedSentence, "？"))
	precedingText = strings.TrimSuffix(precedingText, "?")
	if matched, _ := regexp.MatchString(`\b(吗|是不是|能吗|对吧|好吗|如何|怎么办|行不行|对不对|怎么|为何|为什么|有何|是否|难道)\b`, precedingText); matched {
		return true
	}

	// 检查英文问句
	words := strings.Fields(lowerSentence)
	if len(words) > 0 && englishQuestionWords[words[0]] {
		return true
	}

	return false
}
