package filetracker_test

import (
	"testing"

	"github.com/syndtr/goleveldb/leveldb/opt"

	"github.com/gphotosuploader/gphotos-uploader-cli/internal/datastore/filetracker"
)

func TestLevelDBRepository_Get(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		isErrExpected bool
	}{
		{"Should success", ShouldSuccess, false},
		{"Should fail", ShouldMakeRepoFail, true},
	}

	repo := filetracker.LevelDBRepository{
		DB: &mockedDB{},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.input)
			assertExpectedError(t, tc.isErrExpected, err)
		})
	}
}

func TestLevelDBRepository_Put(t *testing.T) {
	repo := filetracker.LevelDBRepository{
		DB: mockedDB{},
	}
	if err := repo.Put("foo", filetracker.TrackedFile{}); err != nil {
		t.Errorf("error was not expected, err: %s", err)
	}
}

func TestLevelDBRepository_Delete(t *testing.T) {
	repo := filetracker.LevelDBRepository{
		DB: mockedDB{},
	}
	if err := repo.Delete("foo"); err != nil {
		t.Errorf("error was not expected, err: %s", err)
	}
}

func TestLevelDBRepository_Close(t *testing.T) {
	repo := filetracker.LevelDBRepository{
		DB: mockedDB{},
	}
	if err := repo.Close(); err != nil {
		t.Errorf("error was not expected, err: %s", err)
	}
}

type mockedDB struct{}

func (m mockedDB) Get(key []byte, ro *opt.ReadOptions) ([]byte, error) {
	var a []byte
	if string(key) == ShouldMakeRepoFail {
		return a, ErrTestError
	}
	return a, nil
}

func (m mockedDB) Put(key []byte, item []byte, wo *opt.WriteOptions) error {
	return nil
}

func (m mockedDB) Delete(key []byte, wo *opt.WriteOptions) error {
	return nil
}

func (m mockedDB) Close() error {
	return nil
}
