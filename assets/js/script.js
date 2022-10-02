!(function($) {
    'use strict';

    // Mean Menu
    $(".mean-menu").meanmenu({
        meanScreenWidth: "1199",
    });

    // Sticky Header
    $(window).on("scroll", function() {
        var header = $(".sticky-header");
        // If window scroll down .is-sticky class will added to header
        if($(window).scrollTop() >= 200) {
            header.addClass("is-sticky");
        } else {
            header.removeClass("is-sticky");
        }
    });

    // Odometer
    $(".counter").counterUp({
        delay: 10,
        time: 1000
    });

    // Preloader
    $("#preLoader").delay(1000).queue(function() {
        $(this).remove();
    });

    // Testimonial Slider
    $(".testimonial-slider").owlCarousel({
        margin: 40,
        autoplayTimeout: 6500,
        smartSpeed: 500,
        responsiveClass: true,
        responsive: {
            0: {
                items: 1,
            },

            992: {
                items: 2,
            }
        }
    })
    // Feedback Slider
    $(".feedback-slider").owlCarousel({
        margin: 30,
        autoplayTimeout: 6500,
        smartSpeed: 500,
        dots: false,
        loop: true,
        responsiveClass: true,
        responsive: {
            0: {
                items: 1,
            },

            992: {
                items: 2,
            }
        }
    })
    // Feedback Slider 2
    $(".feedback-slider-2").owlCarousel({
        margin: 30,
        autoplayTimeout: 6500,
        smartSpeed: 500,
        dots: false,
        loop: true,
        responsiveClass: true,
        responsive: {
            0: {
                items: 1,
            },

            992: {
                items: 2,
            }
        }
    })
    // Feedback Slider 3
    $(".feedback-slider-3").owlCarousel({
        margin: 30,
        autoplayTimeout: 6500,
        smartSpeed: 500,
        dots: false,
        navContainer: "#feedbackNav",
        navText: [
            "<i class='fal fa-long-arrow-left'></i>",
            "<i class='fal fa-long-arrow-right'></i>"
        ],
        responsiveClass: true,
        responsive: {
            0: {
                items: 1,
            },

            992: {
                items: 2,
            }
        }
    })
    // Feedback Slider 3
    $(".feedback-slider-4").owlCarousel({
        margin: 30,
        autoplayTimeout: 6500,
        smartSpeed: 500,
        dots: false,
        navContainer: "#feedbackNav",
        navText: [
            "<i class='far fa-angle-left'></i>",
            "<i class='far fa-angle-right'></i>"
        ],
        responsiveClass: true,
        responsive: {
            0: {
                items: 1,
            },

            992: {
                items: 2,
            }
        }
    })
    // Sponsor Slider
    $(".sponsor-slider").owlCarousel({
        margin: 30,
        autoplay: true,
        loop: true,
        autoplayTimeout: 3000,
        smartSpeed: 500,
        dots: false,
        nav: false,
        items: 4,
        responsiveClass: true,
        responsive: {
            0: {
                items: 2,
            },

            992: {
                items: 4,
            }
        }
    })

    // Youtube Popup
    $(".youtube-popup").magnificPopup({
        disableOn: 300,
        type: "iframe",
        mainClass: "mfp-fade",
        removalDelay: 160,
        preloader: false,
        fixedContentPos: false
    })

    // Go to Top
    $(window).on("scroll", function() {
        // If window scroll down .active class will added to go-top
        var goTop = $(".go-top");
        if($(window).scrollTop() >= 200) {
            goTop.addClass("active");
        } else {
            goTop.removeClass("active")
        }
    })
    $(".go-top").on("click", function(e) {
        $("html, body").animate({
            scrollTop: 0,
        }, 0 );
    });

    // Countdown Timer
    function makeTimer() {
        var endTime = new Date("May 20, 2024 17:00:00 PDT");
        var endTime = (Date.parse(endTime)) / 1000;
        var now = new Date();
        var now = (Date.parse(now) / 1000);
        var timeLeft = endTime - now;
        var days = Math.floor(timeLeft / 86400);
        var hours = Math.floor((timeLeft - (days * 86400)) / 3600);
        var minutes = Math.floor((timeLeft - (days * 86400) - (hours * 3600)) / 60);
        var seconds = Math.floor((timeLeft - (days * 86400) - (hours * 3600) - (minutes * 60)));
        if (hours < "10") {
            hours = "0" + hours;
        }
        if (minutes < "10") {
            minutes = "0" + minutes;
        }
        if (seconds < "10") {
            seconds = "0" + seconds;
        }
        $("#days .h1").html(days);
        $("#hours .h1").html(hours);
        $("#minutes .h1").html(minutes);
        $("#seconds .h1").html(seconds);
    }
    setInterval(function(){
        makeTimer()
    }, 0);

    // Pricing Toggle Switch
    $("#toggleSwitch").on("change", function(event) {
        if (event.currentTarget.checked) {
            $("#yearly").addClass("active");
            $("#monthly").removeClass("active");
        } else {
            $("#monthly").addClass("active");
            $("#yearly").removeClass("active");
        }
    })

    // Lazy-load Image
    function lazyLoad() {
        window.lazySizesConfig = window.lazySizesConfig || {};
        window.lazySizesConfig.loadMode = 2;
        lazySizesConfig.preloadAfterLoad = true;
    }
    $(document).ready(function() {
        lazyLoad();
    })

})(jQuery);