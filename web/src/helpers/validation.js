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
    yearStart: yup
      .number()
      .required("organisms.createGroupForm.form.yearStart.requiredHelperText"),
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

export function getInviteStudentValidationSchema() {
  return yup.object({
    groupId: yup
      .string()
      .required("organisms.inviteStudentForm.form.groupId.requiredHelperText"),
    email: yup
      .string()
      .email("organisms.inviteStudentValidation.form.email.helperText")
      .required("organisms.inviteStudentForm.form.email.requiredHelperText"),
  });
}

export function getSendEmailValidationSchema() {
  return yup.object({
    subject: yup
      .string()
      .required("organisms.sendEmailForm.form.subject.requiredHelperText"),
    facultyId: yup.string().notRequired(),
    body: yup
      .string()
      .required("organisms.sendEmailForm.form.body.requiredHelperText"),
  });
}

export function getCreateEmailTemplateValidationSchema() {
  return yup.object({
    subject: yup
      .string()
      .required(
        "organisms.createEmailTemplateForm.form.subject.requiredHelperText"
      ),
    body: yup
      .string()
      .required(
        "organisms.createEmailTemplateForm.form.body.requiredHelperText"
      ),
  });
}
