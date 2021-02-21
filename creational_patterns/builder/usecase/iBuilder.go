package usecase

type iBuilder interface {
    setWindowType()
    setDoorType()
    setNumFloor()
    getHouse() House
}

func GetBuilder(builderType string) iBuilder {
    if builderType == "normal" {
        return &normalBuilder{}
    }

    if builderType == "igloo" {
        return &iglooBuilder{}
    }
    return nil
}