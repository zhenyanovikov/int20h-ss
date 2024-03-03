import { Stack, Box, TextField, Button } from "@mui/material";
import { useFormik } from "formik";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";
import { useCreateEmailTemplate } from "../../../api/notification";
import { getCreateEmailTemplateValidationSchema } from "../../../helpers/validation";
import { ROUTE } from "../../../constants/router";

function CreateEmailTemplateForm() {
  const { t } = useTranslation();
  const navigate = useNavigate();

  const {
    mutate: createEmailTemplateMutate,
    isLoading: isCreateEmailTemplateLoading,
  } = useCreateEmailTemplate();

  const formik = useFormik({
    initialValues: {
      subject: "",
      body: "",
    },
    validationSchema: getCreateEmailTemplateValidationSchema(),
    onSubmit: handleSubmit,
  });

  return (
    <form onSubmit={formik.handleSubmit}>
      <Stack spacing={4}>
        <Box sx={{ width: { xs: "100%", md: "60%" } }}>
          <TextField
            fullWidth
            id="subject"
            name="subject"
            label={t("organisms.createEmailTemplateForm.form.subject.label")}
            value={formik.values.subject}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            error={formik.touched.subject && !!formik.errors.subject}
            helperText={formik.touched.subject && t(formik.errors.subject)}
          />
        </Box>
        <Box sx={{ width: { xs: "100%", md: "60%" } }}>
          <TextField
            fullWidth
            id="body"
            name="body"
            label={t("organisms.createEmailTemplateForm.form.body.label")}
            multiline
            rows={8}
            value={formik.values.body}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            error={formik.touched.body && !!formik.errors.body}
            helperText={formik.touched.body && t(formik.errors.body)}
          />
        </Box>
        <Box>
          <Button
            variant="contained"
            type="submit"
            disabled={isCreateEmailTemplateLoading}
          >
            {t("organisms.createEmailTemplateForm.button.create")}
          </Button>
        </Box>
      </Stack>
    </form>
  );

  function handleSubmit(values) {
    createEmailTemplateMutate(values);

    navigate(ROUTE.EMAIL_TEMPLATES);
  }
}

export default CreateEmailTemplateForm;
