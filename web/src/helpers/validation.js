import * as yup from "yup";

export function getInviteTeacherValidationSchema() {
  return yup.object({
    email: yup
      .string()
      .email("organisms.inviteTeacherForm.form.email.helperText")
      .required("organisms.inviteTeacherForm.form.email.requiredHelperText"),
  });
}
