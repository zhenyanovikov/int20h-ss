import * as yup from "yup";

export function getCreateFacultyValidationSchema() {
  return yup.object({
    name: yup
      .string()
      .required("organisms.createFacultyForm.form.name.requiredHelperText"),
  });
}

export function getCreateGroupValidationSchema() {
  return yup.object({
    facultyId: yup
      .string()
      .required("organisms.createGroupForm.form.facultyId.requiredHelperText"),
    name: yup
      .string()
      .required("organisms.createGroupForm.form.name.requiredHelperText"),
  });
}

export function getInviteTeacherValidationSchema() {
  return yup.object({
    email: yup
      .string()
      .email("organisms.inviteTeacherForm.form.email.helperText")
      .required("organisms.inviteTeacherForm.form.email.requiredHelperText"),
  });
}
