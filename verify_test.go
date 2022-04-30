package main

import (
	"encoding/json"
	"testing"
	"time"

	sJWT "github.com/brianvoe/sjwt"
	cristalhqJWT "github.com/cristalhq/jwt"
	jose2go "github.com/dvsekhvalnov/jose2go"
	gbrlsnchsJWT "github.com/gbrlsnchs/jwt/v3"
	goJose "github.com/go-jose/go-jose/v3" // former "gopkg.in/square/go-jose.v2"
	golangJWT "github.com/golang-jwt/jwt/v4"
	karatasJWT "github.com/kataras/jwt"
	jwxJWA "github.com/lestrrat-go/jwx/v2/jwa" // part of "github.com/lestrrat-go/jwx/v2
	jwxJWT "github.com/lestrrat-go/jwx/v2/jwt" // part of "github.com/lestrrat-go/jwx/v2
	pascaldekloeJWT "github.com/pascaldekloe/jwt"
	robbert229JWT "github.com/robbert229/jwt"
	// jwtAuth "github.com/adam-hanna/jwt-auth" // middleware only
	// gojwt "github.com/nickvellios/gojwt" // not actively maintained
	// jwit "github.com/gilbsgilbs/jwit" // uses "gopkg.in/square/go-jose.v2"
)

const testSecret = "sercrethatmaycontainch@r$32chars"

var claims, _ = json.Marshal(map[string]string{"foo": "bar"})

func createTestToken() ([]byte, error) {
	return karatasJWT.Sign(karatasJWT.HS256, []byte(testSecret), claims, karatasJWT.MaxAge(15*time.Minute))
}

// "github.com/cristalhq/jwt"
func BenchmarkVerify_L0(b *testing.B) {
	token, err := createTestToken()
	if err != nil {
		b.Error(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		alg, err := cristalhqJWT.NewHS256([]byte(testSecret))
		if err != nil {
			b.Error(err)
		}
		_, err = cristalhqJWT.ParseAndVerify(token, alg)
		if err != nil {
			b.Error(err)
		}
	}
}

// "github.com/kataras/jwt"
func BenchmarkVerify_L1(b *testing.B) {
	token, err := createTestToken()
	if err != nil {
		b.Error(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err = karatasJWT.Verify(karatasJWT.HS256, []byte(testSecret), token)
		if err != nil {
			b.Error(err)
		}
	}
}

// "github.com/golang-jwt/jwt/v4"
func BenchmarkVerify_L2(b *testing.B) {
	token, err := createTestToken()
	if err != nil {
		b.Error(err)
	}
	tokenString := string(token)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err = golangJWT.Parse(tokenString, func(token *golangJWT.Token) (interface{}, error) {
			return []byte(testSecret), nil
		})
		if err != nil {
			b.Error(err)
		}
	}
}

// "github.com/dvsekhvalnov/jose2go"
func BenchmarkVerify_L3(b *testing.B) {
	token, err := createTestToken()
	if err != nil {
		b.Error(err)
	}
	tokenString := string(token)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, err = jose2go.Decode(tokenString, []byte(testSecret))
		if err != nil {
			b.Error(err)
		}
	}
}

// "github.com/robbert229/jwt"
func BenchmarkVerify_L4(b *testing.B) {
	token, err := createTestToken()
	if err != nil {
		b.Error(err)
	}
	tokenString := string(token)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		algorithm := robbert229JWT.HmacSha256(testSecret)
		_, err = algorithm.DecodeAndValidate(tokenString)
		if err != nil {
			b.Error(err)
		}
	}
}

// "github.com/go-jose/go-jose"
func BenchmarkVerify_L5(b *testing.B) {
	token, err := createTestToken()
	if err != nil {
		b.Error(err)
	}
	tokenString := string(token)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jws, err := goJose.ParseSigned(tokenString)
		if err != nil {
			b.Error(err)
		}
		_, err = jws.Verify([]byte(testSecret))
		if err != nil {
			b.Error(err)
		}
	}
}

// "github.com/lestrrat-go/jwx/v2"
func BenchmarkVerify_L6(b *testing.B) {
	token, err := createTestToken()
	if err != nil {
		b.Error(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err = jwxJWT.Parse(token, jwxJWT.WithKey(jwxJWA.HS256, []byte(testSecret)))
		if err != nil {
			b.Error(err)
		}
	}
}

// "github.com/gbrlsnchs/jwt/v3"
func BenchmarkVerify_L7(b *testing.B) {
	token, err := createTestToken()
	if err != nil {
		b.Error(err)
	}
	var payload = make(map[string]interface{})

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		alg := gbrlsnchsJWT.NewHS256([]byte(testSecret))
		_, err = gbrlsnchsJWT.Verify(token, alg, &payload)
		if err != nil {
			b.Error(err)
		}
	}
}

// "github.com/pascaldekloe/jwt"
func BenchmarkVerify_L8(b *testing.B) {
	token, err := createTestToken()
	if err != nil {
		b.Error(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		alg, err := pascaldekloeJWT.NewHMAC(pascaldekloeJWT.HS256, []byte(testSecret))
		if err != nil {
			b.Error(err)
		}
		_, err = alg.Check(token)
		if err != nil {
			b.Error(err)
		}
	}
}

// "github.com/brianvoe/sjwt"
func BenchmarkVerify_L9(b *testing.B) {
	token, err := createTestToken()
	if err != nil {
		b.Error(err)
	}
	tokenString := string(token)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		valid := sJWT.Verify(tokenString, []byte(testSecret))
		if !valid {
			b.Error("invalid token")
		}
		claims, err := sJWT.Parse(tokenString)
		if err != nil {
			b.Error(err)
		}
		err = claims.Validate()
		if err != nil {
			b.Error(err)
		}
	}
}
