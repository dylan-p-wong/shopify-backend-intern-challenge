package error

import (
  "github.com/go-playground/validator/v10"
)

func GetErrorMsg(fe validator.FieldError) string {
    switch fe.Tag() {
        case "required":
            return fe.Field() + " is required";
        case "lte":
            return fe.Field() + " should be less than " + fe.Param();
        case "gte":
            return fe.Field() + " should be greater than " + fe.Param();
        case "min":
            return fe.Field() + " should be longer than " + fe.Param();
        case "max":
            return fe.Field() + " should be shorter than " + fe.Param();
    }
    return " unknown error";
}
