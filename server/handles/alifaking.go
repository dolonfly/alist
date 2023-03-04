package handles

import (
	"github.com/alist-org/alist/v3/internal/conf"
	"github.com/alist-org/alist/v3/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/url"
	"os"
)

func AliFakingResponse(c *gin.Context) {
	fileName := c.Query("filename")
	//length := c.Query("length")

	if fileName == "!!!!!!.mp4" {
		//common.SuccessResp(c, op.GetPublicSettingsMap())
		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename*=UTF-8''` + url.QueryEscape(fileName),
			//"Content-Type":        "application/oct-stream",
			//"Content-Length":      `` + length,
			"Accept-Ranges": `bytes`,
		}
		c.Header("Content-Disposition", `attachment; filename*=UTF-8''`+url.QueryEscape(fileName))
		//c.Header("Content-Type", "application/oct-stream")
		////c.Header("Content-Length", ``+length)
		//c.Header("Accept-Ranges", `bytes`)
		file, _ := os.Open("/opt/alist/data/default1.mp4")
		fileInfo, _ := os.Lstat("/opt/alist/data/default1.mp4")
		//c.Data(200, "application/oct-stream", []byte{1, 2, 3})

		c.DataFromReader(200, fileInfo.Size(), "application/oct-stream", file, extraHeaders)
		return
	}

	if utils.GetFileType(fileName) == conf.VIDEO {
		//common.SuccessResp(c, op.GetPublicSettingsMap())
		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename*=UTF-8''` + url.QueryEscape(fileName),
			//"Content-Type":        "application/oct-stream",
			//"Content-Length":      `` + length,
			"Accept-Ranges": `bytes`,
		}
		c.Header("Content-Disposition", `attachment; filename*=UTF-8''`+url.QueryEscape(fileName))
		//c.Header("Content-Type", "application/oct-stream")
		////c.Header("Content-Length", ``+length)
		//c.Header("Accept-Ranges", `bytes`)
		file, _ := os.Open("/Users/dll/Downloads/default.mp4")
		fileInfo, _ := os.Lstat("/Users/dll/Downloads/default.mp4")
		//c.Data(200, "application/oct-stream", []byte{1, 2, 3})

		c.DataFromReader(200, fileInfo.Size(), "application/oct-stream", file, extraHeaders)
		return
	} else {
		c.Status(403)
		return
	}

}
