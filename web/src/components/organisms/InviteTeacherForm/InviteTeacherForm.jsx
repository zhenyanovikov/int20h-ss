import { Stack, TextField, Button } from "@mui/material";
import { useTranslation } from "react-i18next";
import { useFormik } from "formik";
import { getInviteTeacherValidationSchema } from "../../../helpers/validation";
import { useInviteTeacher } from "../../../api/teacher";

function InviteTeacherForm() {
  const { t } = useTranslation();
  const { mutate: inviteTeacherMutate, isLoading: isInviteTeacherLoading } =
    useInviteTeacher();

  const formik = useFormik({
    initialValues: {
      email: "",
    },
    validationSchema: getInviteTeacherValidationSchema(),
    onSubmit: handleSubmit,
  });

  return (
    <form onSubmit={formik.handleSubmit}>
      <Stack direction="row" spacing={2}>
        <TextField
          id="email"
          name="email"
          size="small"
          label={t("organisms.inviteTeacherForm.form.email.label")}
          value={formik.values.email}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          error={formik.touched.email && !!formik.errors.email}
          helperText={formik.touched.email && t(formik.errors.email)}
        />
        <Stack>
          <Button
            variant="contained"
            type="submit"
            disabled={isInviteTeacherLoading}
          >
            {t("organisms.inviteTeacherForm.form.button.submit")}
          </Button>
        </Stack>
      </Stack>
    </form>
  );

  function handleSubmit(values, actions) {
    inviteTeacherMutate(values);

    actions.resetForm();
  }
}

export default InviteTeacherForm;
