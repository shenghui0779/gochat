package offia

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/chenghonour/gochat/urls"
	"github.com/chenghonour/gochat/wx"
)

// ImageSize 图片尺寸
type ImageSize struct {
	W int `json:"w"`
	H int `json:"h"`
}

// Position 位置信息
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// ImagePosition 图片位置
type ImagePosition struct {
	LeftTop     Position `json:"left_top"`
	RightTop    Position `json:"right_top"`
	RightBottom Position `json:"right_bottom"`
	LeftBottom  Position `json:"left_bottom"`
}

// CropPosition 裁切位置
type CropPosition struct {
	CropLeft   int `json:"crop_left"`
	CropTop    int `json:"crop_top"`
	CropRight  int `json:"crop_right"`
	CropBottom int `json:"crop_bottom"`
}

// ResultAICrop 图片裁切结果
type ResultAICrop struct {
	Results []*CropPosition `json:"results"`
	ImgSize ImageSize       `json:"img_size"`
}

// AICrop 智能接口 - 图片智能裁切
func AICrop(imgPath string, result *ResultAICrop) wx.Action {
	_, filename := filepath.Split(imgPath)

	return wx.NewPostAction(urls.OffiaAICrop,
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(imgPath))

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFormFile("img", filename, func(w io.Writer) error {
					f, err := os.Open(path)

					if err != nil {
						return err
					}

					defer f.Close()

					if _, err = io.Copy(w, f); err != nil {
						return err
					}

					return nil
				}),
			), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// AICropByURL 智能接口 - 图片智能裁切
func AICropByURL(imgURL string, result *ResultAICrop) wx.Action {
	return wx.NewPostAction(urls.OffiaAICrop,
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// QRCodeScanData 二维码扫描数据
type QRCodeScanData struct {
	TypeName string        `json:"type_name"`
	Data     string        `json:"data"`
	Pos      ImagePosition `json:"pos"`
}

// ResultQRCodeScan 二维码扫描结果
type ResultQRCodeScan struct {
	CodeResults []*QRCodeScanData `json:"code_results"`
	ImgSize     ImageSize         `json:"img_size"`
}

// ScanQRCode 智能接口 - 条码/二维码识别
func ScanQRCode(imgPath string, result *ResultQRCodeScan) wx.Action {
	_, filename := filepath.Split(imgPath)

	return wx.NewPostAction(urls.OffiaScanQRCode,
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(imgPath))

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFormFile("img", filename, func(w io.Writer) error {
					f, err := os.Open(path)

					if err != nil {
						return err
					}

					defer f.Close()

					if _, err = io.Copy(w, f); err != nil {
						return err
					}

					return nil
				}),
			), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// ScanQRCodeByURL 智能接口 - 条码/二维码识别
func ScanQRCodeByURL(imgURL string, result *ResultQRCodeScan) wx.Action {
	return wx.NewPostAction(urls.OffiaScanQRCode,
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// ResultSuperreSolution 图片高清化结果
type ResultSuperreSolution struct {
	MediaID string `json:"media_id"`
}

// SuperreSolution 智能接口 - 图片高清化
func SuperreSolution(imgPath string, result *ResultSuperreSolution) wx.Action {
	_, filename := filepath.Split(imgPath)

	return wx.NewPostAction(urls.OffiaSuperreSolution,
		wx.WithUpload(func() (wx.UploadForm, error) {
			path, err := filepath.Abs(filepath.Clean(imgPath))

			if err != nil {
				return nil, err
			}

			return wx.NewUploadForm(
				wx.WithFormFile("img", filename, func(w io.Writer) error {
					f, err := os.Open(path)

					if err != nil {
						return err
					}

					defer f.Close()

					if _, err = io.Copy(w, f); err != nil {
						return err
					}

					return nil
				}),
			), nil
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// SuperreSolutionByURL 智能接口 - 图片高清化
func SuperreSolutionByURL(imgURL string, result *ResultSuperreSolution) wx.Action {
	return wx.NewPostAction(urls.OffiaSuperreSolution,
		wx.WithQuery("img_url", imgURL),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
