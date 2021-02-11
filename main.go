package main

import (
	"deenya-api/database"
	"deenya-api/handler"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/swaggo/swag"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"

	_ "deenya-api/docs"
)

// @title AvaMed API
// @version 1.0
// @description This is a REST API for AvaMed

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /
func main() {

	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Init()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	corsParams := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(corsParams.Handler)

	handler.TokenAuth = jwtauth.New("HS256", []byte("secret123"), nil)

	r.Route("/consultants", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(handler.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Get("/", handler.MyConsultants)
		})
	})

	r.Route("/clients", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(handler.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Get("/", handler.MyClients)
		})
	})

	r.Route("/users", func(r chi.Router) {
		//paginate?
		r.Get("/search", handler.SearchUsers)
		// Subrouters:
		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/", handler.GetUser)
			r.Get("/posts", handler.UserPosts)
			r.Get("/portfolios", handler.UserPortfolios)
			r.Get("/products", handler.UserProducts)
			r.Get("/threads", handler.UserForumThreads)
			r.Get("/public", handler.UserPublic)
			r.Get("/available", handler.AvailableBookings)
			r.Group(func(r chi.Router) {
				r.Use(jwtauth.Verifier(handler.TokenAuth))
				r.Use(jwtauth.Authenticator)
				r.Get("/me", handler.MyAccount)
				r.Put("/", handler.UpdateUser)
				r.Delete("/", handler.DeleteUser)
			})
		})
	})

	r.Route("/posts", func(r chi.Router) {
		//r.With(handler.PaginatePosts).Get("/", handler.ListPost) // GET /articles
		r.Get("/search", handler.SearchPosts) // GET /articles/search

		// Regexp url parameters:
		r.Get("/{postSlug:[a-z-]+}", handler.SlugPost) // GET /articles/home-is-toronto
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(handler.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Post("/", handler.NewPost)
			r.Get("/", handler.MyPosts) //list own posts
		})

		// Subrouters:
		r.Route("/{postID}", func(r chi.Router) {
			// r.Use(handler.PostCtx)
			r.Get("/", handler.GetPost) // GET /articles/123
			r.Group(func(r chi.Router) {
				r.Use(jwtauth.Verifier(handler.TokenAuth))
				r.Use(jwtauth.Authenticator)
				r.Put("/", handler.UpdatePost)    // PUT /articles/123
				r.Delete("/", handler.DeletePost) // DELETE /articles/123
			})
		})
	})

	r.Route("/bookings", func(r chi.Router) {
		r.Use(jwtauth.Verifier(handler.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/", handler.MyBookings) //list bookings from jwt.user.id and jwt.user.type
		r.Post("/", handler.NewBooking)

		r.Route("/{bookingID}", func(r chi.Router) {
			r.Get("/", handler.GetBooking)
			r.Put("/", handler.UpdateBooking)
			r.Delete("/", handler.DeleteBooking)
			r.Route("/treatment", func(r chi.Router) {
				//TODO-Gor doc
				r.Post("/", handler.NewTreatment)
			})
			//TODO-Gor doc
			r.Route("/questionnaire", func(r chi.Router) {
				r.Post("/", handler.NewQuestionnaire)
			})
		})
	})

	r.Route("/treatments", func(r chi.Router) {
		r.Use(jwtauth.Verifier(handler.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/", handler.MyTreatments)
		r.Route("/{treatmentID}", func(r chi.Router) {
			r.Get("/", handler.GetTreatment)
			r.Put("/", handler.UpdateTreatment)
			r.Delete("/", handler.DeleteTreatment)
		})
	})

	r.Route("/products", func(r chi.Router) {

		r.Post("/", handler.NewProduct)

		r.Route("/models", func(r chi.Router) {
			//r.Post("/", handler.NewProductModel)
			r.Get("/search", handler.SearchProductModels)
			r.Get("/suggest", handler.SuggestProductModels)
			r.Route("/{modelID}", func(r chi.Router) {
				//r.Put("/", handler.UpdateProductModel)
				//r.Get("/", handler.GetProductModel)
			})
		})

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(handler.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Get("/", handler.MyProducts)
		})

		r.Route("/{productID}", func(r chi.Router) {
			r.Get("/", handler.GetProduct)
			r.Group(func(r chi.Router) {
				r.Use(jwtauth.Verifier(handler.TokenAuth))
				r.Use(jwtauth.Authenticator)
				r.Put("/", handler.UpdateProduct)
				r.Delete("/", handler.DeleteProduct)
			})
		})

	})

	r.Route("/categories", func(r chi.Router) {
		r.Get("/", handler.ListCategories)
		r.Get("/{categoryID}", handler.GetCategory)
	})

	r.Route("/portfolios", func(r chi.Router) {

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(handler.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Post("/", handler.NewPortfolio)
			r.Get("/", handler.MyPortfolios)
		})

		r.Route("/{portfolioID}", func(r chi.Router) {
			r.Get("/", handler.GetPortfolio)
			r.Group(func(r chi.Router) {
				r.Use(jwtauth.Verifier(handler.TokenAuth))
				r.Use(jwtauth.Authenticator)
				r.Put("/", handler.UpdatePortfolio)
				r.Delete("/", handler.DeletePortfolio)
			})
		})
	})

	r.Route("/conversations", func(r chi.Router) {
		r.Use(jwtauth.Verifier(handler.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/", handler.MyConversations)
		r.Post("/", handler.NewConversation) //return models.Conversation.ID

		r.Route("/{conversationID}", func(r chi.Router) {
			r.Get("/", handler.GetConversation)
			r.Post("/", handler.NewMessage)
			r.Put("/", handler.UpdateConversation)
			r.Delete("/", handler.DeleteConversation)
			r.Route("/messages/{messageID}", func(r chi.Router) {
				r.Get("/", handler.GetMessage)
				r.Put("/", handler.UpdateMessage)
				r.Delete("/", handler.DeleteMessage)
			})
		})
	})

	r.Route("/media", func(r chi.Router) {

		r.Use(jwtauth.Verifier(handler.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/", handler.MyMedia)
		r.Post("/", handler.NewMedia)

		r.Route("/{mediaID}", func(r chi.Router) {
			r.Get("/", handler.GetMedia)
			r.Put("/", handler.UpdateMedia)
			r.Delete("/", handler.DeleteMedia)
		})
	})

	r.Route("/questionnaires", func(r chi.Router) {
		r.Use(jwtauth.Verifier(handler.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", handler.NewQuestionnaire)

		r.Route("/{questionnaireID}", func(r chi.Router) {
			r.Get("/", handler.GetQuestionnaire)
			r.Post("/", handler.NewQuestion)
			r.Put("/", handler.UpdateQuestionnaire)
			r.Delete("/", handler.DeleteQuestionnaire)
			r.Route("/{questionID}", func(r chi.Router) {
				r.Get("/", handler.GetQuestion)
				r.Put("/", handler.UpdateQuestion)
				r.Delete("/", handler.DeleteQuestion)
			})
		})
	})

	r.Route("/forum", func(r chi.Router) {

		r.Get("/feed", handler.FeedForumThreads)     //universal forum feed
		r.Get("/search", handler.SearchForumThreads) // search forum threads
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(handler.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Get("/", handler.MyForumThreads)
			r.Post("/", handler.NewForumThread)
		})
		r.Route("/{threadID}", func(r chi.Router) {
			r.Get("/", handler.GetForumThread)
			r.Get("/posts", handler.ForumThreadPosts)
			r.Group(func(r chi.Router) {
				r.Use(jwtauth.Verifier(handler.TokenAuth))
				r.Use(jwtauth.Authenticator)
				r.Post("/", handler.NewForumPost)
				r.Put("/", handler.UpdateForumThread)
				r.Delete("/", handler.DeleteForumThread)
			})
			r.Route("/{postID}", func(r chi.Router) {
				r.Get("/", handler.GetForumPost)
				r.Put("/", handler.UpdateForumPost)
				r.Delete("/", handler.DeleteForumPost)
			})
		})
	})

	r.Route("/clinics", func(r chi.Router) {

		//r.Get("/search", handler.SearchClinics)

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(handler.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Post("/", handler.NewClinic)
			r.Get("/", handler.MyClinic)
		})

		r.Route("/{clinicID}", func(r chi.Router) {
			r.Get("/", handler.GetClinic)
			r.Get("/consultants", handler.ListClinicConsultants)
			r.Group(func(r chi.Router) {
				r.Use(jwtauth.Verifier(handler.TokenAuth))
				r.Use(jwtauth.Authenticator)
				r.Get("/requests", handler.ListClinicRequests)
				r.Post("/", handler.NewClinicRequest)
				r.Put("/", handler.UpdateClinic)
				r.Delete("/", handler.LeaveClinic)
			})

			r.Route("/{consultantID}", func(r chi.Router) {
				r.Use(jwtauth.Verifier(handler.TokenAuth))
				r.Use(jwtauth.Authenticator)
				r.Get("/accept", handler.AcceptClinicRequest)
				r.Get("/remove", handler.RemoveClinicMember)
			})
		})
	})

	r.Route("/finance", func(r chi.Router) {
		r.Use(jwtauth.Verifier(handler.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/invoices", handler.MyInvoices)
		r.Post("/invoices", handler.NewInvoice)

		r.Route("/{invoiceID}", func(r chi.Router) {
			r.Get("/", handler.GetInvoice)
			r.Put("/", handler.UpdateInvoice)
			// r.Post("/pay", handler.PayInvoice)
		})

		r.Route("/stripe", func(r chi.Router) {
			r.Post("/connect", handler.NewConnectAccount)
			r.Post("/", handler.NewCustomer)

			r.Route("/{customerID}", func(r chi.Router) {
				r.Get("/", handler.GetCustomer)

				r.Put("/", handler.UpdateCustomer)
				r.Delete("/", handler.DeleteCustomer)
			})
			// r.Route("/connect", func(r chi.Router) {
			// 	r.Route("{connectID}", func(r chi.Router) {

			// 	})
			// 	r.Get("/", handler.GetStripeConnect)
			// 	r.Get("/", handler.NewStripeConnect)
			// 	r.Get("/", handler.UpdateStripeConnect)
			// 	r.Get("/", handler.GetStripeConnect)
			// 	r.Route("/payout", func(r chi.Router) {
			// 		r.Get("/", handler.GetStripeConnectPayout)
			// 		r.Get("/", handler.NewStripeConnectPayout)
			// 		r.Get("/", handler.UpdateStripeConnectPayout)
			// 		r.Get("/", handler.DeleteStripeConnectPayout)
			// 	})
			// })
		})
		//add card
		//update card
		//get list of cards
		//add bank account
		//update bank
		//integrate stripe
		r.Route("/geo", func(r chi.Router) {
			// r.Route("countries", handler.GetCountriesAndRegions)
		})
	})

	r.Route("/search", func(r chi.Router) {
		//paginate?
		r.Get("/consultants", handler.SearchConsultants) //query, distance
		r.Get("/posts", handler.SearchPosts)
		r.Get("/portfolios", handler.SearchPortfolios)
		r.Get("/products", handler.SearchProducts)
		r.Get("/clinics", handler.SearchClinics)
	})

	r.Post("/login", handler.Login)
	r.Post("/register", handler.Register)

	r.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		str, err := swag.ReadDoc()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handler.WriteAsJSON(w, []byte(str))
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}

}
