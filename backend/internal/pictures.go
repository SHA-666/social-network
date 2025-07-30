package internal

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)
func SavePics(file multipart.File, fileHeader *multipart.FileHeader, defaultPath string) (string, error) {
	if fileHeader == nil {
		return defaultPath, nil
	}
	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	// Dossier racine où les images sont stockées pour être servies
	relDir := "../internal/assets/uploads/profile_pic"
	if err := os.MkdirAll(relDir, os.ModePerm); err != nil {
		return "", err
	}

	// Emplacement complet pour l'enregistrement du fichier
	dstPath := filepath.Join(relDir, filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	// Chemin relatif que tu peux stocker en base
	return filepath.Join("uploads/profile_pic", filename), nil
}
func SavePost(file multipart.File, fileHeader *multipart.FileHeader, defaultPath string) (string, error) {
	if fileHeader == nil {
		return defaultPath, nil
	}
	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	// Dossier racine où les images sont stockées pour être servies
	relDir := "../internal/assets/uploads/posts"
	if err := os.MkdirAll(relDir, os.ModePerm); err != nil {
		return "", err
	}

	// Emplacement complet pour l'enregistrement du fichier
	dstPath := filepath.Join(relDir, filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	// Chemin relatif que tu peux stocker en base
	return filepath.Join("uploads/posts", filename), nil
}