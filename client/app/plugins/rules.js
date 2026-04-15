import { defineRule } from "vee-validate";

export default defineNuxtPlugin((vueApp) => {
  (defineRule("onerequired", (value, [], ctx) => {
    if (!value) {
      return "At least one selection required";
    }
    if (typeof value == "object") {
      return value?.length || "At least one selection required";
    } else {
      return !!value || "At least one selection required";
    }
  }),
    defineRule("Selectrequired", (value) => {
      if (!value || (Array.isArray(value) && value.length === 0)) {
        return "Field is required"; // Error message when the field is empty or an array is empty
      }
      return true; // No error
    }));

  (defineRule("arrayRequired", (value, [], ctx) => {
    return value?.length || "At least one required";
  }),
    defineRule("required", (value, [], ctx) => {
      return !!value || value?.length || "Field Required";
    }),
    defineRule("boolReq", (value, [], ctx) => {
      return typeof value == "boolean" || "Field Required";
    }));
  (defineRule("fayda", (value, [], ctx) => {
    return !value || value.length == 10 || "Fayda ID must be 10 Digits";
  }),
    defineRule("verify_fayda", (value, [], ctx) => {
      return value || "Verify your Id";
    }),
    defineRule("array_object_required", (value, [], ctx) => {
      return value?.length || "Field Required";
    }),
    defineRule("yes_or_no", (value, [], ctx) => {
      return value?.length || "Field Required";
    }),
    defineRule("dyarrayrequired", (value, [], ctx) => {
      return (
        ctx.rule.params[0]?.includes("No_savings") ||
        value?.length ||
        "Field Required"
      );
    }),
    defineRule("lenInterval", (value, [], context) => {
      if (!value) {
        return true;
      }
      if (
        value?.length >= context.rule.params[0] &&
        value?.length <= context.rule.params[1]
      ) {
        return true;
      } else {
        let show =
          context.rule.params[0] == context.rule.params[1]
            ? `must be ${context.rule.params[0]} digits`
            : `must be between ${context.rule.params[0]} - ${context.rule.params[1]} digits`;
        return `${context.field} ${show}`;
      }
    }));
  defineRule("url", (value, context) => {
    return (
      !value ||
      /^((http|https):\/\/)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$/.test(
        value,
      ) ||
      "Not valid url"
    );
  });
  (defineRule("dyrequired", (value, [], ctx) => {
    return (
      ctx.rule.params[0] == "no" ||
      ctx.rule.params[0] == "No" ||
      ctx.rule.params[0] == undefined ||
      ctx.rule.params[0] == "sole_proprietorship" ||
      (typeof ctx.rule.params[0] == "boolean" && ctx.rule.params[0] == false) ||
      !!value ||
      value?.length ||
      "Field Required"
    );
  }),
    defineRule("arrya_object_dyrequired", (value, [], ctx) => {
      return (
        ctx.rule.params[0] == "no" ||
        ctx.rule.params[0] == "No" ||
        ctx.rule.params[0] == undefined ||
        value?.length ||
        "Field Required"
      );
    }));
  (defineRule("dyrequirednot", (value, [], ctx) => {
    return (
      ctx.rule.params[0] == "yes" ||
      ctx.rule.params[0] == "Yes" ||
      ctx.rule.params[0] == undefined ||
      ctx.rule.params[0] == "sole_proprietorship" ||
      !!value ||
      value?.length ||
      "Field Required"
    );
  }),
    defineRule("number", (value) => {
      return !value || /^[0-9]+$/.test(value) || "Number only";
    }),
    defineRule("numberFromZero", (value, [], ctx) => {
      return /^(?:0|[1-9]\d*)$/.test(value) || "Number only";
    }));
  (defineRule("numrange", (value) => {
    return !value || /^[0-9-]+$/.test(value) || "Number only";
  }),
    defineRule("email", (value) => {
      return (
        !value ||
        /[A-Za-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?/.test(
          value,
        ) ||
        "Not Valid Email"
      );
    }),
    defineRule("International_phone_number", (value) => {
      return (
        !value ||
        /^\+(?:[0-9] ?){6,14}[0-9]$/.test(value) ||
        "Not valid phone number"
      );
    }),
    defineRule("minLength", (value, [], context) => {
      if (value.length >= context.rule.params[0]) {
        return true;
      } else {
        return `${context.field} is must be greater than ${context.rule.params[0]}`;
      }
    }));

  (defineRule("ethiopian_phone_number", (value) => {
    return !value || /^\+251[79]\d{8}$/.test(value) || "Not valid phone number";
  }),
    defineRule("ethio_phone", (value) => {
      return !value || /^(7|9)\d{8}$/.test(value) || "Not valid phone number";
    }),
    defineRule("password", (value) => {
      return !value || value.length >= 8 || "Must be greater than 8";
    }),
    defineRule("confirmed", (value, [other]) => {
      return !value || value === other || "Password is not the same";
    }));
});

(defineRule("dyrequiredemail", (value, [other], ctx) => {
  return (
    ctx.rule.params[0] != "email" ||
    !!value ||
    value?.length ||
    "Field Required"
  );
}),
  defineRule("dyrequiredlink", (value, [other], ctx) => {
    return (
      ctx.rule.params[0] != "link" ||
      !!value ||
      value?.length ||
      "Field Required"
    );
  }),
  defineRule("validPassword", (value) => {
    const strongRegex = new RegExp(
      "^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*])(?=.{8,})",
    );
    return (
      !value ||
      strongRegex.test(value) ||
      "The password should contain letters, numbers, uppercase and special symbols."
    );
  }));

defineRule("amChar", (value) => {
  return (
    !value ||
    /[\u1200-\u137F]/.test(value.replace(/\s/g)) ||
    "This filed only accepts Amharic characters"
  );
});

defineRule("enChar", (value) => {
  return (
    !value ||
    !/[\u1200-\u137F]/.test(value.replace(/\s/g)) ||
    "This filed only accepts English characters"
  );
});

defineRule("checkMinimumInMax", (value, [other]) => {
  if (other) {
    return (
      !value ||
      parseInt(value) >= parseInt(other) ||
      i18n.global.t("max_and_min_check")
    );
  }
  return true;
});

defineRule("ageAtLeast15", (value) => {
  if (!value) {
    return true; // Allow empty values if the field is optional
  }
  const birthYear = new Date(value).getFullYear();
  const currentYear = new Date().getFullYear();
  const age = currentYear - birthYear;
  if (age >= 15) {
    return true; // Valid if the year is 2000 or earlier
  }

  return "Birth year must be at least 15 years ago"; // Error message
});
