package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Companies struct {
	CMGUnmaskedID                 string
	CMGUnmaskedName               string
	ClientTier                    string
	GCPStream                     string
	GCPBusiness                   string
	CMGGlobalBU                   string
	CMGSegmentName                string
	GlobalControlPoint            string
	GCPGeography                  string
	GlobalRelationshipManagerName string
	REVENUE_FY14                  string
	REVENUE_FY15                  string
	Deposits_EOP_FY14             string
	Deposits_EOP_FY15x            string
	TotalLimits_EOP_FY14          string
	TotalLimits_EOP_FY15          string
	TotalLimits_EOP_FY15x         string
	RWAFY15                       string
	RWAFY14                       string
	REVRWA_FY14                   string
	REVRWA_FY15                   string
	NPAT_AllocEq_FY14             string
	NPAT_AllocEq_FY15X            string
	Company_Avg_Activity_FY14     string
	Company_Avg_Activity_FY15     string
	ROE_FY14                      string
	ROE_FY15                      string
}

const filePath = "./companies.csv"

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/", getHome)
	router.POST("/", postHome)

	router.GET("/detail/:id", getDetail)
	router.POST("/update", postUpdate)

	router.Run()
}

func getHome(c *gin.Context) {
	csvFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var comp Companies
	var companies []Companies

	for _, each := range csvData[1:] {
		comp.CMGUnmaskedID = each[0]
		// if index == 0 {
		// 	comp.CMGUnmaskedID = "ini 1"
		// }
		// fmt.Println(each[0])
		comp.CMGUnmaskedName = each[1]
		comp.ClientTier = each[2]
		comp.GCPStream = each[3]
		comp.GCPBusiness = each[4]
		comp.CMGGlobalBU = each[5]
		comp.CMGSegmentName = each[6]
		comp.GlobalControlPoint = each[7]
		comp.GCPGeography = each[8]
		comp.GlobalRelationshipManagerName = each[9]
		comp.REVENUE_FY14 = each[10]
		comp.REVENUE_FY15 = each[11]
		comp.Deposits_EOP_FY14 = each[12]
		comp.Deposits_EOP_FY15x = each[13]
		comp.TotalLimits_EOP_FY14 = each[14]
		comp.TotalLimits_EOP_FY15 = each[15]
		comp.TotalLimits_EOP_FY15x = each[16]
		comp.RWAFY15 = each[17]
		comp.RWAFY14 = each[18]
		comp.REVRWA_FY14 = each[19]
		comp.REVRWA_FY15 = each[20]
		comp.NPAT_AllocEq_FY14 = each[21]
		comp.NPAT_AllocEq_FY15X = each[22]
		comp.Company_Avg_Activity_FY14 = each[23]
		comp.Company_Avg_Activity_FY15 = each[24]
		comp.ROE_FY14 = each[25]
		comp.ROE_FY15 = each[26]
		companies = append(companies, comp)
	}

	// Convert to JSON
	_, err = json.Marshal(companies)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Println(string(jsonData))

	// jsonFile, err := os.Create("./data.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer jsonFile.Close()

	// jsonFile.Write(jsonData)
	// jsonFile.Close()

	c.JSON(200, gin.H{
		"data": companies,
	})
}

func postHome(c *gin.Context) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	CMGUnmaskedID := c.PostForm("CMGUnmaskedID")
	CMGUnmaskedName := c.PostForm("CMGUnmaskedName")
	ClientTier := c.PostForm("ClientTier")
	GCPStream := c.PostForm("GCPStream")
	GCPBusiness := c.PostForm("GCPBusiness")
	CMGGlobalBU := c.PostForm("CMGGlobalBU")
	CMGSegmentName := c.PostForm("CMGSegmentName")
	GlobalControlPoint := c.PostForm("GlobalControlPoint")
	GCPGeography := c.PostForm("GCPGeography")
	GlobalRelationshipManagerName := c.PostForm("GlobalRelationshipManagerName")
	REVENUE_FY14 := c.PostForm("REVENUE_FY14")
	REVENUE_FY15 := c.PostForm("REVENUE_FY15")
	Deposits_EOP_FY14 := c.PostForm("Deposits_EOP_FY14")
	Deposits_EOP_FY15x := c.PostForm("Deposits_EOP_FY15x")
	TotalLimits_EOP_FY14 := c.PostForm("TotalLimits_EOP_FY14")
	TotalLimits_EOP_FY15 := c.PostForm("TotalLimits_EOP_FY15")
	TotalLimits_EOP_FY15x := c.PostForm("TotalLimits_EOP_FY15x")
	RWAFY15 := c.PostForm("RWAFY15")
	RWAFY14 := c.PostForm("RWAFY14")
	REVRWA_FY14 := c.PostForm("REVRWA_FY14")
	REVRWA_FY15 := c.PostForm("REVRWA_FY15")
	NPAT_AllocEq_FY14 := c.PostForm("NPAT_AllocEq_FY14")
	NPAT_AllocEq_FY15X := c.PostForm("NPAT_AllocEq_FY15X")
	Company_Avg_Activity_FY14 := c.PostForm("Company_Avg_Activity_FY14")
	Company_Avg_Activity_FY15 := c.PostForm("Company_Avg_Activity_FY15")
	ROE_FY14 := c.PostForm("ROE_FY14")
	ROE_FY15 := c.PostForm("ROE_FY15")

	var data [][]string
	data = append(data, []string{
		CMGUnmaskedID,
		CMGUnmaskedName,
		ClientTier,
		GCPStream,
		GCPBusiness,
		CMGGlobalBU,
		CMGSegmentName,
		GlobalControlPoint,
		GCPGeography,
		GlobalRelationshipManagerName,
		REVENUE_FY14,
		REVENUE_FY15,
		Deposits_EOP_FY14,
		Deposits_EOP_FY15x,
		TotalLimits_EOP_FY14,
		TotalLimits_EOP_FY15,
		TotalLimits_EOP_FY15x,
		RWAFY15,
		RWAFY14,
		REVRWA_FY14,
		REVRWA_FY15,
		NPAT_AllocEq_FY14,
		NPAT_AllocEq_FY15X,
		Company_Avg_Activity_FY14,
		Company_Avg_Activity_FY15,
		ROE_FY14,
		ROE_FY15,
	})

	w := csv.NewWriter(f)
	w.WriteAll(data)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"data": map[string]string{
			"CMGUnmaskedID":                 CMGUnmaskedID,
			"CMGUnmaskedName":               CMGUnmaskedName,
			"ClientTier":                    ClientTier,
			"GCPStream":                     GCPStream,
			"GCPBusiness":                   GCPBusiness,
			"CMGGlobalBU":                   CMGGlobalBU,
			"CMGSegmentName":                CMGSegmentName,
			"GlobalControlPoint":            GlobalControlPoint,
			"GCPGeography":                  GCPGeography,
			"GlobalRelationshipManagerName": GlobalRelationshipManagerName,
			"REVENUE_FY14":                  REVENUE_FY14,
			"REVENUE_FY15":                  REVENUE_FY15,
			"Deposits_EOP_FY14":             Deposits_EOP_FY14,
			"Deposits_EOP_FY15x":            Deposits_EOP_FY15x,
			"TotalLimits_EOP_FY14":          TotalLimits_EOP_FY14,
			"TotalLimits_EOP_FY15":          TotalLimits_EOP_FY15,
			"TotalLimits_EOP_FY15x":         TotalLimits_EOP_FY15x,
			"RWAFY15":                       RWAFY15,
			"RWAFY14":                       RWAFY14,
			"REVRWA_FY14":                   REVRWA_FY14,
			"REVRWA_FY15":                   REVRWA_FY15,
			"NPAT_AllocEq_FY14":             NPAT_AllocEq_FY14,
			"NPAT_AllocEq_FY15X":            NPAT_AllocEq_FY15X,
			"Company_Avg_Activity_FY14":     Company_Avg_Activity_FY14,
			"Company_Avg_Activity_FY15":     Company_Avg_Activity_FY15,
			"ROE_FY14":                      ROE_FY14,
			"ROE_FY15":                      ROE_FY15,
		},
	})
}
func getDetail(c *gin.Context) {
	csvFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// var comp Companies
	var companies []Companies
	var detail Companies

	for _, each := range csvData {
		if each[0] == c.Param("id") {
			detail.CMGUnmaskedID = each[0]
			detail.CMGUnmaskedName = each[1]
			detail.ClientTier = each[2]
			detail.GCPStream = each[3]
			detail.GCPBusiness = each[4]
			detail.CMGGlobalBU = each[5]
			detail.CMGSegmentName = each[6]
			detail.GlobalControlPoint = each[7]
			detail.GCPGeography = each[8]
			detail.GlobalRelationshipManagerName = each[9]
			detail.REVENUE_FY14 = each[10]
			detail.REVENUE_FY15 = each[11]
			detail.Deposits_EOP_FY14 = each[12]
			detail.Deposits_EOP_FY15x = each[13]
			detail.TotalLimits_EOP_FY14 = each[14]
			detail.TotalLimits_EOP_FY15 = each[15]
			detail.TotalLimits_EOP_FY15x = each[16]
			detail.RWAFY15 = each[17]
			detail.RWAFY14 = each[18]
			detail.REVRWA_FY14 = each[19]
			detail.REVRWA_FY15 = each[20]
			detail.NPAT_AllocEq_FY14 = each[21]
			detail.NPAT_AllocEq_FY15X = each[22]
			detail.Company_Avg_Activity_FY14 = each[23]
			detail.Company_Avg_Activity_FY15 = each[24]
			detail.ROE_FY14 = each[25]
			detail.ROE_FY15 = each[26]
			companies = append(companies, detail)
		}
	}

	// Convert to JSON
	_, err = json.Marshal(companies)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Println(string(jsonData))

	// jsonFile, err := os.Create("./data.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer jsonFile.Close()

	// jsonFile.Write(jsonData)
	// jsonFile.Close()

	c.JSON(200, gin.H{
		"data": companies,
	})
}

func postUpdate(c *gin.Context) {

	//open file
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := csv.NewReader(f).ReadAll()

	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	CMGUnmaskedID := c.PostForm("CMGUnmaskedID")
	CMGUnmaskedName := c.PostForm("CMGUnmaskedName")
	ClientTier := c.PostForm("ClientTier")
	GCPStream := c.PostForm("GCPStream")
	GCPBusiness := c.PostForm("GCPBusiness")
	CMGGlobalBU := c.PostForm("CMGGlobalBU")
	CMGSegmentName := c.PostForm("CMGSegmentName")
	GlobalControlPoint := c.PostForm("GlobalControlPoint")
	GCPGeography := c.PostForm("GCPGeography")
	GlobalRelationshipManagerName := c.PostForm("GlobalRelationshipManagerName")
	REVENUE_FY14 := c.PostForm("REVENUE_FY14")
	REVENUE_FY15 := c.PostForm("REVENUE_FY15")
	Deposits_EOP_FY14 := c.PostForm("Deposits_EOP_FY14")
	Deposits_EOP_FY15x := c.PostForm("Deposits_EOP_FY15x")
	TotalLimits_EOP_FY14 := c.PostForm("TotalLimits_EOP_FY14")
	TotalLimits_EOP_FY15 := c.PostForm("TotalLimits_EOP_FY15")
	TotalLimits_EOP_FY15x := c.PostForm("TotalLimits_EOP_FY15x")
	RWAFY15 := c.PostForm("RWAFY15")
	RWAFY14 := c.PostForm("RWAFY14")
	REVRWA_FY14 := c.PostForm("REVRWA_FY14")
	REVRWA_FY15 := c.PostForm("REVRWA_FY15")
	NPAT_AllocEq_FY14 := c.PostForm("NPAT_AllocEq_FY14")
	NPAT_AllocEq_FY15X := c.PostForm("NPAT_AllocEq_FY15X")
	Company_Avg_Activity_FY14 := c.PostForm("Company_Avg_Activity_FY14")
	Company_Avg_Activity_FY15 := c.PostForm("Company_Avg_Activity_FY15")
	ROE_FY14 := c.PostForm("ROE_FY14")
	ROE_FY15 := c.PostForm("ROE_FY15")

	// loop data
	var id bool = false
	for i := 1; i < len(rows); i++ {
		if rows[i][0] == c.PostForm("CMGUnmaskedID") {
			// fmt.Println(rows[i][1])
			// rows[i][0] = CMGUnmaskedID
			rows[i][1] = CMGUnmaskedName
			rows[i][2] = ClientTier
			rows[i][3] = GCPStream
			rows[i][4] = GCPBusiness
			rows[i][5] = CMGGlobalBU
			rows[i][6] = CMGSegmentName
			rows[i][7] = GlobalControlPoint
			rows[i][8] = GCPGeography
			rows[i][9] = GlobalRelationshipManagerName
			rows[i][10] = REVENUE_FY14
			rows[i][11] = REVENUE_FY15
			rows[i][12] = Deposits_EOP_FY14
			rows[i][13] = Deposits_EOP_FY15x
			rows[i][14] = TotalLimits_EOP_FY14
			rows[i][15] = TotalLimits_EOP_FY15
			rows[i][16] = TotalLimits_EOP_FY15x
			rows[i][17] = RWAFY15
			rows[i][18] = RWAFY14
			rows[i][19] = REVRWA_FY14
			rows[i][20] = REVRWA_FY15
			rows[i][21] = NPAT_AllocEq_FY14
			rows[i][22] = NPAT_AllocEq_FY15X
			rows[i][23] = Company_Avg_Activity_FY14
			rows[i][24] = Company_Avg_Activity_FY15
			rows[i][25] = ROE_FY14
			rows[i][26] = ROE_FY15
			id = true
		}
		// rows[i] = append(rows[i], sum(rows[i]))
	}

	// create file again
	f, err = os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = csv.NewWriter(f).WriteAll(rows)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	if id {
		c.JSON(200, gin.H{
			"msg": "berhasil update id " + CMGUnmaskedID,
		})
	} else {
		c.JSON(500, gin.H{
			"msg": "id not found",
		})
	}
}

// libarary udpate

func sum(row []string) string {
	sum := 0
	for _, s := range row {
		x, err := strconv.Atoi(s)
		if err != nil {
			return "NA"
		}
		sum += x
	}
	return strconv.Itoa(sum)
}
