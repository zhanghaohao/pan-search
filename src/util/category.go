package util

func ExtToCategory(ext string) string {
	if (ext == "mp4" || ext == "MP4" || ext == "avi" || ext == "webm" || ext == "3gp" || ext == "wmv" || ext == "mkv" || ext == "mpg" || ext == "vod" || ext == "vob" || ext == "mov" || ext == "flv" || ext == "swf" || ext == "rm" || ext == "rmvb") {
		return "video"
	} else if (ext == "doc" || ext == "pdf" || ext == "docx" || ext == "txt" || ext == "ppt" || ext == "pptx") {
		return "document"
	} else if (ext == "mp3" || ext == "wav" || ext == "wma" || ext == "mid") {
		return "audio"
	} else if (ext == "torrent") {
		return "seed"
	} else if (ext == "jpg" || ext == "svg" || ext == "jpeg" || ext == "gif" || ext == "bmp" || ext == "png" || ext == "jpe" || ext == "cur" || ext == "svgz" || ext == "ico" || ext == "heic" || ext == "heif" || ext == "avci" || ext == "webp" || ext == "tif") {
		return "picture"
	} else if (ext == "") {
		return "folder"
	} else if (ext == "zip" || ext == "rar" || ext == "7z" || ext == "cab" || ext == "tar" || ext == "arc" || ext == "pcf" || ext == "exe") {
		return "archive"
	} else {
		return "unknown"
	}
}
