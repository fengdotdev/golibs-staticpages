package helperfuncs

import googleuuid "github.com/google/uuid"

// use uuid v4 to generate a unique id randomly
func GenerateId() string {
	uuidV4 := googleuuid.New()
	return uuidV4.String()
}

// this ensures that uuid from google is encapsulated in case of future changes
type SomeUUID struct {
	uuid string
}

func NewSomeUUID() SomeUUID {
	return SomeUUID{
		uuid: GenerateId(),
	}
}

func (s *SomeUUID) String() string {
	return s.uuid
}

func GenerateIdDeterministic(name string, namespace SomeUUID) string {
	namespacePrime := googleuuid.MustParse(namespace.String())
	uuidV5 := googleuuid.NewMD5(namespacePrime, []byte(name))
	return uuidV5.String()
}

func GenerateIdDeterministicWithHASH(name string, toHash string) string {
	hashed := HashMD5([]byte(toHash))
	namespacePrime := googleuuid.MustParse(hashed)
	uuidV5 := googleuuid.NewMD5(namespacePrime, []byte(name))
	return uuidV5.String()
}
