import {
  Stack,
  Box,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  TextField,
  Button,
} from "@mui/material";
import { useLocation, useNavigate } from "react-router-dom";
import { useFormik } from "formik";
import { useTranslation } from "react-i18next";
import { getSendEmailValidationSchema } from "../../../helpers/validation";
import { useGetFaculties } from "../../../api/faculty";
import { useSendEmail } from "../../../api/notification";
import { ROUTE } from "../../../constants/router";

function SendEmailForm() {
  const { t } = useTranslation();
  const navigte = useNavigate();
  const location = useLocation();

  const emailTemplate = location.state?.emailTemplate;

  const { mutate: sendEmailMutate, isLoading: isSendEmailLoading } =
    useSendEmail();
  const { data: facultiesData } = useGetFaculties();

  const formik = useFormik({
    initialValues: {
      subject: emailTemplate ? emailTemplate.subject : "",
      facultyId: "",
      body: emailTemplate ? emailTemplate.body : "",
    },
    validationSchema: getSendEmailValidationSchema(),
    onSubmit: handleSubmit,
  });

  const faculties = facultiesData || [];

  return (
    <form onSubmit={formik.handleSubmit}>
      <Stack spacing={4}>
        <Box sx={{ width: { xs: "100%", md: "60%" } }}>
          <TextField
            fullWidth
            id="subject"
            name="subject"
            label={t("organisms.sendEmailForm.form.subject.label")}
            value={formik.values.subject}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            error={formik.touched.subject && !!formik.errors.subject}
            helperText={formik.touched.subject && t(formik.errors.subject)}
          />
        </Box>
        <Box sx={{ width: { xs: "100%", md: "30%" } }}>
          <FormControl fullWidth>
            <InputLabel id="facultyId">
              {t("organisms.sendEmailForm.form.facultyId.label")}
            </InputLabel>
            <Select
              labelId="facultyId"
              id="facultyId"
              name="facultyId"
              value={formik.values.facultyId}
              label={t("organisms.sendEmailForm.form.facultyId.label")}
              onChange={formik.handleChange}
            >
              {faculties.map((faculty) => (
                <MenuItem key={faculty.id} value={faculty.id}>
                  {faculty.name}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        </Box>
        <Box sx={{ width: { xs: "100%", md: "60%" } }}>
          <TextField
            fullWidth
            id="body"
            name="body"
            label={t("organisms.sendEmailForm.form.body.label")}
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
            disabled={isSendEmailLoading}
          >
            {t("organisms.sendEmailForm.button.send")}
          </Button>
        </Box>
      </Stack>
    </form>
  );

  function handleSubmit(values) {
    sendEmailMutate(values);

    navigte(ROUTE.EMAIL_TEMPLATES);
  }
}

export default SendEmailForm;
