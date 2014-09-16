package py.com.codelab.android.fraserochuck;

import android.app.Activity;
import android.os.AsyncTask;
import android.os.Bundle;
import android.util.JsonReader;
import android.util.Log;
import android.view.Menu;
import android.view.MenuItem;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;
import android.widget.Toast;

import com.example.david.papa.R;

import org.apache.http.HttpEntity;
import org.apache.http.HttpResponse;
import org.apache.http.client.HttpClient;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.entity.StringEntity;
import org.apache.http.impl.client.DefaultHttpClient;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;


public class ChuckMainActivity extends Activity {

    private static final String BASE_ADDR = "http://192.168.43.184:8081/frases/";
    TextView tv;
    EditText edFrase;
    EditText edUsuario;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_my);
        final Button btnOtra = (Button)findViewById(R.id.btnOtra);
        tv = (TextView)findViewById(R.id.textView);
        btnOtra.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                btnOtra.setPressed(true);
                new CheckFraseTask().execute(BASE_ADDR);
            }
        });
    }


    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.my, menu);
        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        // Handle action bar item clicks here. The action bar will
        // automatically handle clicks on the Home/Up button, so long
        // as you specify a parent activity in AndroidManifest.xml.
        int id = item.getItemId();
        if (id == R.id.add_frase) {
            setContentView(R.layout.add_frase);
            final Button btnAgregar = (Button)findViewById(R.id.btnAgregar);
            btnAgregar.setOnClickListener(new View.OnClickListener() {
                @Override
                public void onClick(View v) {
                    btnAgregar.setPressed(true);
                    new AgregarFraseTask().execute(BASE_ADDR + "/add");
                }
            });
        }
        return super.onOptionsItemSelected(item);
    }

    HttpClient httpclient = new DefaultHttpClient();

    private class CheckFraseTask extends AsyncTask<String, Integer, String> {

        protected String doInBackground(String... urls) {

            HttpGet httpget = new HttpGet(urls[0]);

            HttpResponse response;

            String result = "";

            try {
                response = httpclient.execute(httpget);

                // Examine the response status
                Log.i("Conectado", response.getStatusLine().toString());

                // Get hold of the response entity
                HttpEntity entity = response.getEntity();
                // If the response does not enclose an entity, there is no need
                // to worry about connection release

                if (entity != null) {

                    // A Simple JSON Response Read
                    InputStream instream = entity.getContent();

                    JsonReader reader = new JsonReader(new InputStreamReader(instream, "UTF-8"));

                    reader.beginObject();
                    while (reader.hasNext()) {
                        String name = reader.nextName();
                        if (name.equals("LaFrase")) {
                            result = reader.nextString();
                        } else {
                            reader.skipValue();
                        }
                    }
                    reader.endObject();

                    // now you have the string representation of the HTML request
                    instream.close();
                }


            } catch (Exception e) {
                Log.e("Error:", e.getMessage(), e);
            }

            return result;
        }

        // This is called when doInBackground() is finished
        protected void onPostExecute(String result) {
            //showNotification("Downloaded " + result + " bytes");
            Log.i("Resultado:", result);
            tv.setText(result);
        }

    }


    private class AgregarFraseTask extends AsyncTask<String, Integer, String> {

        protected String doInBackground(String... urls) {

            HttpPost httppost = new HttpPost(urls[0]);

            edFrase = (EditText)findViewById(R.id.txtAddFrase);
            edUsuario = (EditText)findViewById(R.id.txtUsuario);

            try {

                httppost.setEntity(new StringEntity("{" +
                        "\"LaFrase\":\"" + edFrase.getText() + "\"," +
                        "\"Usuario\":\"" + edUsuario.getText() + "\"," +
                        "\"Ip\":\"notimpementedyet\"," +
                        "}"));

            } catch (Exception e) {
                e.printStackTrace();
                return "Error al enviar datos: " + e.getMessage();
            }

            HttpResponse response;

            String result = "";

            try {
                response = httpclient.execute(httppost);

                Log.i("Conectado", response.getStatusLine().toString());

                HttpEntity entity = response.getEntity();

                if (entity != null) {

                    InputStream instream = entity.getContent();

                    result = convertStreamToString(instream);

                    instream.close();
                }


            } catch (Exception e) {
                Log.e("Error:", e.getMessage(), e);
            }

            return result;
        }

        protected void onPostExecute(String result) {
            Log.i("Resultado:", result);
            if(result.contains("Se inserto correctamente"))
                setContentView(R.layout.activity_my);
            else
                Toast.makeText(getApplicationContext(), "Error al insertar la frase. Vuelva a intentarlo por favor.", Toast.LENGTH_LONG);
        }

    }


    private String convertStreamToString(InputStream is) {
        /*
         * To convert the InputStream to String we use the BufferedReader.readLine()
         * method. We iterate until the BufferedReader return null which means
         * there's no more data to read. Each line will appended to a StringBuilder
         * and returned as String.
         */
            BufferedReader reader = new BufferedReader(new InputStreamReader(is));
            StringBuilder sb = new StringBuilder();

            String line = null;
            try {
                while ((line = reader.readLine()) != null) {
                    sb.append(line + "\n");
                }
            } catch (IOException e) {
                e.printStackTrace();
            } finally {
                try {
                    is.close();
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
            return sb.toString();
        }

}
